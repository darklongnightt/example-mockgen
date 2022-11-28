//go:generate mockgen -source=service.go -destination=./mocks/service.go
package user

import (
	"example-mockgen/models"
)

// https://github.com/golang/go/wiki/CodeReviewComments#interfaces
// > interface should belong in the package that uses it value
// so based on this quote, the interface should exclusive to the current package, hence, unexported.

type iRepository interface {
	Insert(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

type iS3Client interface {
	UploadFile(string) error
}

// Service contains dependencies and Service methods for users
type Service struct {
	repo iRepository
	s3   iS3Client
}

// New inits new user Service with interfaces (that can be real or mock implementations)
func New(repo iRepository, s3 iS3Client) *Service {
	return &Service{
		repo: repo,
		s3:   s3,
	}
}

// CreateUser creates a user in db and do something
func (u *Service) CreateUser(user *models.User) (*models.User, error) {
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
func (u *Service) UpdateUser(user *models.User) (*models.User, error) {
	resp, err := u.repo.Update(user)
	if err != nil {
		return nil, err
	}

	// Do something - rest of the API Service

	return resp, nil
}
