package repositories

import (
	"github.com/butlerhq/butler/internal/services/auth/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo create a new repo for models.User
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) WithTransaction(db *gorm.DB) *UserRepo {
	return NewUserRepo(db)
}

// HashPassword return a hash for a given password
func (u *UserRepo) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

// VerifyPassword return an error if a user HashPassword doesn't match password
func (u *UserRepo) VerifyPassword(user *models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
}

// CreateOne and save it in database
func (repo *UserRepo) CreateOne(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return &models.User{}, err
	}
	return user, nil
}

// ListByIDs list users using a list of IDs
func (repo *UserRepo) ListByIDs(userIDs []string) ([]models.User, error) {
	users := []models.User{}
	if err := repo.db.Find(&users, userIDs).Error; err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// FindByID a user in database
func (repo *UserRepo) FindByID(userID string) (*models.User, error) {
	user := &models.User{}
	if err := repo.db.Model(user).Where("id = ?", userID).Take(user).Error; err != nil {
		return &models.User{}, err
	}
	return user, nil
}

// FindByEmail a user in database
func (repo *UserRepo) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := repo.db.Model(user).Where("email = ?", email).Take(user).Error; err != nil {
		return &models.User{}, err
	}
	return user, nil
}

// UpdateOne
func (repo *UserRepo) UpdateOne(user *models.User) error {
	if err := repo.db.Model(user).Updates(user).Error; err != nil {
		return err
	}
	return nil
}
