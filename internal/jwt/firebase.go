package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// CertsAPIEndpoint is endpoint of getting Public Key.
const CertsAPIEndpoint = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

type FirebaseJWTConfig struct {
	CertsEndpoint string
}

type FirebaseJWTManager struct {
	config *FirebaseJWTConfig
}

func NewFirebaseJWTManager() *FirebaseJWTManager {
	return &FirebaseJWTManager{
		config: &FirebaseJWTConfig{
			CertsEndpoint: CertsAPIEndpoint,
		},
	}
}

func (m *FirebaseJWTManager) getCertificates() (certs map[string]string, err error) {
	res, err := http.Get(CertsAPIEndpoint)
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	json.Unmarshal(data, &certs)
	return
}

// GetCertificate returns certificate.
func (m *FirebaseJWTManager) getCertificate(kid string) (cert []byte, err error) {
	certs, err := m.getCertificates()
	if err != nil {
		return
	}
	certString := certs[kid]
	cert = []byte(certString)
	err = nil
	return
}

// GetCertificateFromToken returns cert from token.
func (m *FirebaseJWTManager) GetCertificateFromToken(token *jwt.Token) ([]byte, error) {
	// Get kid
	kid, ok := token.Header["kid"]
	if !ok {
		return []byte{}, errors.New("kid not found")
	}
	kidString, ok := kid.(string)
	if !ok {
		return []byte{}, errors.New("kid cast error to string")
	}
	return m.getCertificate(kidString)
}

// Verify the token payload.
func (m *FirebaseJWTManager) verifyPayload(t *jwt.Token, projectID string) (ok bool, uid string) {
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return
	}
	// Verify User
	claimsAud, ok := claims["aud"].(string)
	if claimsAud != projectID || !ok {
		return
	}
	// Verify issued at
	iss := fmt.Sprintf("https://securetoken.google.com/%s", projectID)
	claimsIss, ok := claims["iss"].(string)
	if claimsIss != iss || !ok {
		return
	}
	// sub is uid of user.
	uid, ok = claims["sub"].(string)
	if !ok {
		return
	}
	return
}

func (m *FirebaseJWTManager) readPublicKey(cert []byte) (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode(cert)
	if publicKeyBlock == nil {
		return nil, errors.New("invalid public key data")
	}
	if publicKeyBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("invalid public key type: %s", publicKeyBlock.Type)
	}
	c, err := x509.ParseCertificate(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}
	return publicKey, nil
}

func (m *FirebaseJWTManager) VerifyJWT(t, projectID string) (uid string, ok bool) {
	parsed, _ := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		cert, err := m.GetCertificateFromToken(t)
		if err != nil {
			return "", err
		}
		publicKey, err := m.readPublicKey(cert)
		if err != nil {
			return "", err
		}
		return publicKey, nil
	})

	ok = parsed.Valid
	if !ok {
		return
	}
	// Verify header.
	if parsed.Header["alg"] != "RS256" {
		ok = false
		return
	}
	// Verify payload.
	ok, uid = m.verifyPayload(parsed, projectID)
	return
}
