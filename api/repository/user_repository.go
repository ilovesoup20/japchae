package repository

import (
	"github.com/ilovesoup20/japchae/model"
)

// UserRepository .
type UserRepository interface {
	Create(user model.User) error
	ReadByID(userID int) (model.User, error)
	Update(user model.User) error
	Delete(userID int) error
	FindAll() ([]model.User, error)
}

// UserRepositoryImpl .
type UserRepositoryImpl struct {
}

func (r *UserRepositoryImpl) Create(user model.User) error {

}

func (r *UserRepositoryImpl) ReadByID(userID int) (model.User, error) {

}

func (r *UserRepositoryImpl) Update(user model.User) error {

}

func (r *UserRepositoryImpl) Delete(userID int) error {

}

func (r *UserRepositoryImpl) FindAll() ([]model.User, error) {

}
