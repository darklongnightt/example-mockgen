package services_test

import (
	"errors"
	"example-mockgen/models"
	"example-mockgen/services"
	mocks "example-mockgen/services/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Simple(t *testing.T) {
	ctr := gomock.NewController(t)

	s3 := mocks.NewMockIS3Client(ctr)
	s3.EXPECT().UploadFile(gomock.Any()).Return(nil).Times(1)

	repo := mocks.NewMockIUserRepo(ctr)
	repo.EXPECT().Insert(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)

	h := services.NewUserService(repo, s3)
	got, err := h.CreateUser(&models.User{})
	assert.Equal(t, &models.User{Name: "mock user"}, got)
	assert.Equal(t, nil, err)
}

func TestCreateUser_MultipleScenarios(t *testing.T) {
	ctr := gomock.NewController(t)

	tests := []struct {
		name string
		want *models.User
		repo services.IUserRepo // User DB interface to be mocked
		s3   services.IS3Client // S3 client interface to be mocked
		err  error
	}{
		{
			name: "create a user successfully",
			repo: func() services.IUserRepo {
				mockRepo := mocks.NewMockIUserRepo(ctr)
				mockRepo.EXPECT().Insert(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)
				return mockRepo
			}(),
			s3: func() services.IS3Client {
				mockS3 := mocks.NewMockIS3Client(ctr)
				mockS3.EXPECT().UploadFile(gomock.Any()).Return(nil).Times(1)
				return mockS3
			}(),
			want: &models.User{Name: "mock user"},
			err:  nil,
		},
		{
			name: "db returns error",
			repo: func() services.IUserRepo {
				mockRepo := mocks.NewMockIUserRepo(ctr)
				mockRepo.EXPECT().Insert(gomock.Any()).Return(nil, errors.New("db error")).Times(1)
				return mockRepo
			}(),
			want: nil,
			err:  errors.New("db error"),
		},
		{
			name: "s3 returns error",
			repo: func() services.IUserRepo {
				mockRepo := mocks.NewMockIUserRepo(ctr)
				mockRepo.EXPECT().Insert(gomock.Any()).Return(&models.User{Name: "mock user"}, nil).Times(1)
				return mockRepo
			}(),
			s3: func() services.IS3Client {
				mockS3 := mocks.NewMockIS3Client(ctr)
				mockS3.EXPECT().UploadFile(gomock.Any()).Return(errors.New("s3 error")).Times(1)
				return mockS3
			}(),
			want: nil,
			err:  errors.New("s3 error"),
		},
	}

	for _, tc := range tests {
		h := services.NewUserService(tc.repo, tc.s3)
		got, err := h.CreateUser(nil)
		assert.Equal(t, tc.want, got, tc.name)
		assert.Equal(t, tc.err, err, tc.name)
	}
}
