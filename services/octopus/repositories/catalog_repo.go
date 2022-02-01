package repositories

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/butlerhq/butler/internal/logger"
	"github.com/butlerhq/butler/services/octopus/connectors"
	"github.com/butlerhq/butler/services/octopus/models"
)

type CatalogRepo struct {
	mu            sync.Mutex
	connectorsMap map[string]models.CatalogConnector
}

func NewCatalogRepo(cfg *connectors.Config) *CatalogRepo {
	googleConnector := connectors.NewGoogleConnector(cfg.Google, cfg.RedirectURL)
	linkedinConnector := connectors.NewLinkedinConnector(cfg.Linkedin, cfg.RedirectURL)
	return &CatalogRepo{
		connectorsMap: map[string]models.CatalogConnector{
			googleConnector.Name():   googleConnector,
			linkedinConnector.Name(): linkedinConnector,
		},
	}
}

// ListAvailableConnectors returns a list of available connector
func (repo *CatalogRepo) ListAvailableConnectors() []models.CatalogConnector {
	tmp := make([]models.CatalogConnector, 0, len(repo.connectorsMap))
	repo.mu.Lock()
	for _, value := range repo.connectorsMap {
		fmt.Printf("Conn name: %s", value.Name())
		tmp = append(tmp, value)
	}
	repo.mu.Unlock()
	return tmp
}

// GetConnector returns a list of available connector
func (repo *CatalogRepo) GetConnector(name string) (models.CatalogConnector, bool) {
	repo.mu.Lock()
	connector, ok := repo.connectorsMap[name]
	repo.mu.Unlock()
	return connector, ok
}

func (repo *CatalogRepo) ExchangeOAuthCode(ctx context.Context, provider, code string) (*oauth2.Token, error) {
	logger.Debug(ctx, "About to exchange oauth code", zap.String("provider", provider))
	repo.mu.Lock()
	connector, ok := repo.connectorsMap[provider]
	fmt.Printf("%v", connector)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unknown connector: %s"))
	}
	repo.mu.Unlock()

	token, err := connector.ExchangeCode(ctx, code)
	if err != nil {
		logger.Error(ctx, "Unable to exchange oauth code", zap.Error(err), zap.String("provider", connector.Name()))
		return nil, status.Errorf(codes.InvalidArgument, "Unable to exchange oauth code")
	}
	return token, nil
}
