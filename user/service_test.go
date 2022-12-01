package user_test

import (
	"errors"
	"example-mockgen/models"
	"example-mockgen/user"
	mocks "example-mockgen/user/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// A simple test case that covers only success scenario
// To demo how mocks work
func TestAddUser_Simple(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish() // Not required for go1.14+

	s3 := mocks.NewMockiS3Client(ctr)
	s3.EXPECT().UploadFile(gomock.Any()).Return(nil).Times(1)

	repo := mocks.NewMockiRepository(ctr)
	repo.EXPECT().AddUser(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)

	s := user.New(repo, s3)
	got, err := s.AddUser(&models.User{})
	assert.Equal(t, &models.User{Name: "mock user"}, got)
	assert.Equal(t, nil, err)
}

// Using testify test suite that provide hooks for before and after tests
// Table testing to cover multiple scenarios
type UserSuite struct {
	suite.Suite

	controller   *gomock.Controller
	mockRepo     *mocks.MockiRepository
	mockS3Client *mocks.MockiS3Client
	service      *user.Service
}

// SetupTest runs before all tests to init testing dependencies
func (s *UserSuite) SetupTest() {
	s.controller = gomock.NewController(s.T())
	s.mockRepo = mocks.NewMockiRepository(s.controller)
	s.mockS3Client = mocks.NewMockiS3Client(s.controller)
	s.service = user.New(s.mockRepo, s.mockS3Client)
}

// SetupTest runs after each test for cleanups
func (s *UserSuite) TearDownTest() {
	s.controller.Finish() // Not required for go1.14+
}

// TestUserSuite is required to for go test to run all tests in a suite
func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) TestAddUser() {
	tests := []struct {
		name      string
		initMocks func()
		input     *models.User
		want      *models.User
		err       error
	}{
		{
			name: "create user success",
			initMocks: func() {
				s.mockRepo.EXPECT().AddUser(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)
				s.mockS3Client.EXPECT().UploadFile(gomock.Any()).Return(nil).Times(1)
			},
			input: &models.User{Name: "input user"},
			want:  &models.User{Name: "mock user"},
			err:   nil,
		},
		{
			name: "db error",
			initMocks: func() {
				s.mockRepo.EXPECT().AddUser(gomock.Any()).Return(nil, errors.New("db error")).Times(1)
			},
			input: &models.User{Name: "input user"},
			want:  nil,
			err:   errors.New("db error"),
		},
		{
			name: "s3 error",
			initMocks: func() {
				s.mockRepo.EXPECT().AddUser(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)
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
			got, err := s.service.AddUser(tc.input)
			assert.Equal(s.T(), tc.want, got, tc.name)
			assert.Equal(s.T(), tc.err, err, tc.name)
		})
	}
}

func (s *UserSuite) TestGetUsers() {
	tests := []struct {
		name      string
		want      []*models.User
		initMocks func()
		err       error
	}{
		{
			name: "update user success",
			initMocks: func() {
				s.mockRepo.EXPECT().GetUsers().Return([]*models.User{
					{
						ID:   "123",
						Name: "mock user",
						Age:  12,
					},
				}, nil).Times(1)
			},
			want: []*models.User{
				{
					ID:   "123",
					Name: "mock user",
					Age:  12,
				},
			},
			err: nil,
		},
		{
			name: "db error",
			initMocks: func() {
				s.mockRepo.EXPECT().GetUsers().Return(nil, errors.New("db error")).Times(1)
			},
			want: nil,
			err:  errors.New("db error"),
		},
	}

	for _, tc := range tests {
		s.Run(tc.name, func() {
			tc.initMocks()
			got, err := s.service.GetUsers()
			assert.Equal(s.T(), tc.want, got, tc.name)
			assert.Equal(s.T(), tc.err, err, tc.name)
		})
	}
}
