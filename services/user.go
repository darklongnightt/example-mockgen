//go:generate mockgen -source=user.go -destination=./mocks/user_mocks.go
package services

import (
	"example-mockgen/models"
)

// IUserRepo repo interface used by this Service
type IUserRepo interface {
	Insert(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

// IS3Client interface with methods to connect to S3
type IS3Client interface {
	UploadFile(string) error
}

// UserService contains dependencies and Service methods for users
type UserService struct {
	repo IUserRepo
	s3   IS3Client
}

// NewUserService inits new user Service with interfaces (that can be real or mock implementations)
func NewUserService(repo IUserRepo, s3 IS3Client) *UserService {
	return &UserService{
		repo: repo,
		s3:   s3,
	}
}

// CreateUser creates a user in db and do something
func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	resp, err := u.repo.Insert(user)
	if err != nil {
		return nil, err
	}

	if err := u.s3.UploadFile(resp.Name); err != nil {
		return nil, err
	}

	// Do something - rest of the API Service

	return resp, nil
}

// CreateUser udpates a user in db and do something
func (u *UserService) UpdateUser(user *models.User) (*models.User, error) {
	resp, err := u.repo.Update(user)
	if err != nil {
		return nil, err
	}

	// Do something - rest of the API Service

	return resp, nil
}
