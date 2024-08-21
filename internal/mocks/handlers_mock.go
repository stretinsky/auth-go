// Code generated by MockGen. DO NOT EDIT.
// Source: handlers.go
//
// Generated by this command:
//
//	mockgen -source=handlers.go -destination=../../../mocks/handlers_mock.go -package mock
//

// Package mock is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/bogatyr285/auth-go/internal/auth/entity"
	jwt "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(ctx context.Context, username string) (entity.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", ctx, username)
	ret0, _ := ret[0].(entity.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), ctx, username)
}

// RegisterUser mocks base method.
func (m *MockUserRepository) RegisterUser(ctx context.Context, u entity.UserAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserRepositoryMockRecorder) RegisterUser(ctx, u any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserRepository)(nil).RegisterUser), ctx, u)
}

// MockCryptoPassword is a mock of CryptoPassword interface.
type MockCryptoPassword struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoPasswordMockRecorder
}

// MockCryptoPasswordMockRecorder is the mock recorder for MockCryptoPassword.
type MockCryptoPasswordMockRecorder struct {
	mock *MockCryptoPassword
}

// NewMockCryptoPassword creates a new mock instance.
func NewMockCryptoPassword(ctrl *gomock.Controller) *MockCryptoPassword {
	mock := &MockCryptoPassword{ctrl: ctrl}
	mock.recorder = &MockCryptoPasswordMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoPassword) EXPECT() *MockCryptoPasswordMockRecorder {
	return m.recorder
}

// ComparePasswords mocks base method.
func (m *MockCryptoPassword) ComparePasswords(fromUser, fromDB string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePasswords", fromUser, fromDB)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ComparePasswords indicates an expected call of ComparePasswords.
func (mr *MockCryptoPasswordMockRecorder) ComparePasswords(fromUser, fromDB any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePasswords", reflect.TypeOf((*MockCryptoPassword)(nil).ComparePasswords), fromUser, fromDB)
}

// HashPassword mocks base method.
func (m *MockCryptoPassword) HashPassword(password string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockCryptoPasswordMockRecorder) HashPassword(password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockCryptoPassword)(nil).HashPassword), password)
}

// MockJWTManager is a mock of JWTManager interface.
type MockJWTManager struct {
	ctrl     *gomock.Controller
	recorder *MockJWTManagerMockRecorder
}

// MockJWTManagerMockRecorder is the mock recorder for MockJWTManager.
type MockJWTManagerMockRecorder struct {
	mock *MockJWTManager
}

// NewMockJWTManager creates a new mock instance.
func NewMockJWTManager(ctrl *gomock.Controller) *MockJWTManager {
	mock := &MockJWTManager{ctrl: ctrl}
	mock.recorder = &MockJWTManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWTManager) EXPECT() *MockJWTManagerMockRecorder {
	return m.recorder
}

// IssueToken mocks base method.
func (m *MockJWTManager) IssueToken(userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueToken", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueToken indicates an expected call of IssueToken.
func (mr *MockJWTManagerMockRecorder) IssueToken(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueToken", reflect.TypeOf((*MockJWTManager)(nil).IssueToken), userID)
}

// VerifyToken mocks base method.
func (m *MockJWTManager) VerifyToken(tokenString string) (*jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", tokenString)
	ret0, _ := ret[0].(*jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockJWTManagerMockRecorder) VerifyToken(tokenString any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockJWTManager)(nil).VerifyToken), tokenString)
}