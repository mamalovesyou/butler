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

func NewCatalogRepo(cfg *connectors.ConnectorsConfig) *CatalogRepo {
	googleConnector := connectors.NewGoogleConnector(cfg.Google)
	linkedinConnector := connectors.NewLinkedinConnector(cfg.Linkedin)
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

func (repo *CatalogRepo) ExchangeOAuthCode(ctx context.Context, provider, code string) (*oauth2.Token, error) {
	repo.mu.Lock()
	connector, ok := repo.connectorsMap[provider]
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
