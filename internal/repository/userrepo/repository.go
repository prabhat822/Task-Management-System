package userrepo

import (
    "errors"
	"time"

	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db:        db,
		tableName: "users",
	}
}

func (ur *userRepository) Create(user *entity.User) error {
	userModel := UserModel{}.FromUserEntity(*user)
	timeNow := time.Now()
	userModel.CreatedAt = &timeNow
	err := ur.db.Table(ur.tableName).Create(&userModel).Error
	return err
}

func (ur *userRepository) GetByEmail(email string) (entity.User, error) {
	userModel := UserModel{}
	err := ur.db.Table(ur.tableName).Where("email = ?", email).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, repository.ErrUserNotFound
		}
		return entity.User{}, err
	}
	return *userModel.ToUserEntity(), nil
}

func (ur *userRepository) GetByID(id string) (entity.User, error) {
	return entity.User{}, nil
}
