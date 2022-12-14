// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	models "example-mockgen/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockiRepository is a mock of iRepository interface.
type MockiRepository struct {
	ctrl     *gomock.Controller
	recorder *MockiRepositoryMockRecorder
}

// MockiRepositoryMockRecorder is the mock recorder for MockiRepository.
type MockiRepositoryMockRecorder struct {
	mock *MockiRepository
}

// NewMockiRepository creates a new mock instance.
func NewMockiRepository(ctrl *gomock.Controller) *MockiRepository {
	mock := &MockiRepository{ctrl: ctrl}
	mock.recorder = &MockiRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiRepository) EXPECT() *MockiRepositoryMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockiRepository) AddUser(user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockiRepositoryMockRecorder) AddUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockiRepository)(nil).AddUser), user)
}

// GetUsers mocks base method.
func (m *MockiRepository) GetUsers() ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockiRepositoryMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockiRepository)(nil).GetUsers))
}

// MockiS3Client is a mock of iS3Client interface.
type MockiS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockiS3ClientMockRecorder
}

// MockiS3ClientMockRecorder is the mock recorder for MockiS3Client.
type MockiS3ClientMockRecorder struct {
	mock *MockiS3Client
}

// NewMockiS3Client creates a new mock instance.
func NewMockiS3Client(ctrl *gomock.Controller) *MockiS3Client {
	mock := &MockiS3Client{ctrl: ctrl}
	mock.recorder = &MockiS3ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiS3Client) EXPECT() *MockiS3ClientMockRecorder {
	return m.recorder
}

// UploadFile mocks base method.
func (m *MockiS3Client) UploadFile(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockiS3ClientMockRecorder) UploadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockiS3Client)(nil).UploadFile), arg0)
}
