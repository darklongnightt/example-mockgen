package services_test

import (
	"errors"
	"example-mockgen/models"
	"example-mockgen/services"
	mocks "example-mockgen/services/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

/*
A simple test case that covers happy path to usage of gomocks
*/
func TestCreateUser_Simple(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish() // Not required for go1.14+

	s3 := mocks.NewMockIS3Client(ctr)
	s3.EXPECT().UploadFile(gomock.Any()).Return(nil).Times(1)

	repo := mocks.NewMockIUserRepo(ctr)
	repo.EXPECT().Insert(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)

	h := services.NewUserService(repo, s3)
	got, err := h.CreateUser(&models.User{})
	assert.Equal(t, &models.User{Name: "mock user"}, got)
	assert.Equal(t, nil, err)
}

/*
Using testify test suite that provide hooks for before and after tests
Table testing to cover all scenarios
*/
type UserSuite struct {
	suite.Suite

	controller   *gomock.Controller
	mockUserRepo *mocks.MockIUserRepo
	mockS3Client *mocks.MockIS3Client
	service      *services.UserService
}

// SetupTest runs before all tests to init testing dependencies
func (s *UserSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.mockUserRepo = mocks.NewMockIUserRepo(s.controller)
	s.mockS3Client = mocks.NewMockIS3Client(s.controller)
	s.service = services.NewUserService(s.mockUserRepo, s.mockS3Client)
}

// SetupTest runs after each test for cleanups
func (s *UserSuite) TearDownTest() {
	s.controller.Finish() // Not required for go1.14+
}

// TestUserSuite is required to for go test to run all tests in a suite
func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) TestCreateUser() {
	tests := []struct {
		name      string
		initMocks func()
		input     *models.User
		want      *models.User
		err       error
	}{
		{
			name: "create a user successfully",
			initMocks: func() {
				s.mockUserRepo.EXPECT().Insert(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)
				s.mockS3Client.EXPECT().UploadFile(gomock.Any()).Return(nil).Times(1)
			},
			input: &models.User{Name: "input user"},
			want:  &models.User{Name: "mock user"},
			err:   nil,
		},
		{
			name: "db returns error",
			initMocks: func() {
				s.mockUserRepo.EXPECT().Insert(gomock.Any()).Return(nil, errors.New("db error")).Times(1)
			},
			input: &models.User{Name: "input user"},
			want:  nil,
			err:   errors.New("db error"),
		},
		{
			name: "s3 returns error",
			initMocks: func() {
				s.mockUserRepo.EXPECT().Insert(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)
				s.mockS3Client.EXPECT().UploadFile(gomock.Any()).Return(errors.New("s3 error")).Times(1)
			},
			input: &models.User{Name: "input user"},
			want:  nil,
			err:   errors.New("s3 error"),
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			tc.initMocks()
			got, err := s.service.CreateUser(tc.input)
			assert.Equal(s.T(), tc.want, got, tc.name)
			assert.Equal(s.T(), tc.err, err, tc.name)
		})
	}
}

func (s *UserSuite) TestUpdateUser() {
	tests := []struct {
		name      string
		input     *models.User
		want      *models.User
		initMocks func()
		err       error
	}{
		{
			name: "update a user successfully",
			initMocks: func() {
				s.mockUserRepo.EXPECT().Update(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)
			},
			input: &models.User{Name: "input user"},
			want:  &models.User{Name: "mock user"},
			err:   nil,
		},
		{
			name: "db returns error",
			initMocks: func() {
				s.mockUserRepo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("db error")).Times(1)
			},
			input: &models.User{Name: "input user"},
			want:  nil,
			err:   errors.New("db error"),
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			tc.initMocks()
			got, err := s.service.UpdateUser(tc.input)
			assert.Equal(s.T(), tc.want, got, tc.name)
			assert.Equal(s.T(), tc.err, err, tc.name)
		})
	}
}
