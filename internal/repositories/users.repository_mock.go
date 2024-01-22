package repositories

import (
	"coffee-shop-golang/internal/models"

	"github.com/stretchr/testify/mock"
)

type UsersRepositoryMock struct {
	mock.Mock
}

// RepositoryCreateUsers implements IUsersRepository.
func (urm *UsersRepositoryMock) RepositoryCreateUsers(body *models.UsersModel, hashedPassword string) (int, error) {
	args := urm.Mock.Called(body)
	return args.Get(0).(int), args.Error(1)
}

// RepositoryDeleteUsers implements IUsersRepository.
func (urm *UsersRepositoryMock) RepositoryDeleteUsers(id string) (int, error) {
	args := urm.Mock.Called(id)
	return args.Get(0).(int), args.Error(1)
}

// RepositoryGetAllUsers implements IUsersRepository.
func (urm *UsersRepositoryMock) RepositoryGetAllUsers(name string, page string, limit string, sort string) ([]models.UsersResponseModel, error) {
	args := urm.Mock.Called(name, page, limit, sort)
	return args.Get(0).([]models.UsersResponseModel), args.Error(1)
}

// RepositoryUpdateImgUsers implements IUsersRepository.
func (urm *UsersRepositoryMock) RepositoryUpdateImgUsers(usersImage string, id string) error {
	args := urm.Mock.Called(usersImage, id)
	return args.Error(0)
}

// RepositoryUpdateUsers implements IUsersRepository.
func (urm *UsersRepositoryMock) RepositoryUpdateUsers(body *models.UpdateUserModel, hashedPassword string, id string) error {
	args := urm.Mock.Called(body, hashedPassword, id)
	return args.Error(1)
}

// RepositoryUsersById implements IUsersRepository.
func (urm *UsersRepositoryMock) RepositoryUsersById(id string) ([]models.UsersResponseModel, error) {
	args := urm.Mock.Called(id)
	return args.Get(0).([]models.UsersResponseModel), args.Error(1)
}

// RepositoryCountUsers implements IUsersRepository.
func (urm *UsersRepositoryMock) RepositoryCountUsers(name string) ([]string, error) {
	args := urm.Mock.Called(name)
	return args.Get(0).([]string), args.Error(1)
}

// RepositoryDeleteUsers implements IUsersRepository.
// func (urm *UsersRepositoryMock) RepositoryDeleteUsers(id string) (int64, error) {
// 	args := urm.Mock.Called(id)
// 	return args.Get(0).(int64), args.Error(1)
// }


// RepsitoryCreateUsers implements IUsersRepository.
// func (urm *UsersRepositoryMock) RepsitoryCreateUsers(body *models.UsersModel, hashedPassword string) (int, error) {
// 	args := urm.Mock.Called(body, hashedPassword)
// 	return args.Get(0).(int), args.Error(1)
// }

// RepsitoryGetAllUsers implements IUsersRepository.
func (urm *UsersRepositoryMock) RepsitoryGetAllUsers(page string, limit string, name string, sort string) ([]models.UsersModel, error) {
	args := urm.Mock.Called(page, limit, name, sort)
	return args.Get(0).([]models.UsersModel), args.Error(1)
}

// RepsitoryUpdateUsers implements IUsersRepository.
// func (urm *UsersRepositoryMock) RepsitoryUpdateUsers(body *models.UpdateUsersModel, id string) error {
// 	args := urm.Mock.Called(body, id)
// 	return args.Error(0)
// }
