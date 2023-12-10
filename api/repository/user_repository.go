package repository

import (
	"github.com/ilovesoup20/japchae/model"
	"gorm.io/gorm"
)

// UserRepository .
type UserRepository interface {
	Create(user *model.User) error
	// ReadByID(userID int) (model.User, error)
	// Update(user model.User) error
	// Delete(userID int) error
	FindAll() ([]model.User, error)
}

// UserRepositoryImpl .
type UserRepositoryImpl struct {
	DB *gorm.DB
}

// Create .
func (r *UserRepositoryImpl) Create(user *model.User) error {
	// user := model.User{Username: "charles", Email: "charles@ilovesoup.co"}
	result := r.DB.Create(user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// func (r *UserRepositoryImpl) ReadByID(userID int) (model.User, error) {

// }

// func (r *UserRepositoryImpl) Update(user model.User) error {

// }

// func (r *UserRepositoryImpl) Delete(userID int) error {

// }

func (r *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User
	r.DB.Find(&users)
	return users, nil
}
