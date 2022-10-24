// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/core/ports/ports.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	context "context"
	game "minesweeper-go/internal/core/domain/game"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGamesRepository is a mock of GamesRepository interface.
type MockGamesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGamesRepositoryMockRecorder
}

// MockGamesRepositoryMockRecorder is the mock recorder for MockGamesRepository.
type MockGamesRepositoryMockRecorder struct {
	mock *MockGamesRepository
}

// NewMockGamesRepository creates a new mock instance.
func NewMockGamesRepository(ctrl *gomock.Controller) *MockGamesRepository {
	mock := &MockGamesRepository{ctrl: ctrl}
	mock.recorder = &MockGamesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGamesRepository) EXPECT() *MockGamesRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockGamesRepository) Get(arg0 context.Context, arg1 string) (game.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(game.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGamesRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGamesRepository)(nil).Get), arg0, arg1)
}

// Save mocks base method.
func (m *MockGamesRepository) Save(arg0 context.Context, arg1 game.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockGamesRepositoryMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockGamesRepository)(nil).Save), arg0, arg1)
}

// MockGamesService is a mock of GamesService interface.
type MockGamesService struct {
	ctrl     *gomock.Controller
	recorder *MockGamesServiceMockRecorder
}

// MockGamesServiceMockRecorder is the mock recorder for MockGamesService.
type MockGamesServiceMockRecorder struct {
	mock *MockGamesService
}

// NewMockGamesService creates a new mock instance.
func NewMockGamesService(ctrl *gomock.Controller) *MockGamesService {
	mock := &MockGamesService{ctrl: ctrl}
	mock.recorder = &MockGamesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGamesService) EXPECT() *MockGamesServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockGamesService) Create(ctx context.Context, name string, height, width, bombs int) (game.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, name, height, width, bombs)
	ret0, _ := ret[0].(game.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockGamesServiceMockRecorder) Create(ctx, name, height, width, bombs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGamesService)(nil).Create), ctx, name, height, width, bombs)
}

// Get mocks base method.
func (m *MockGamesService) Get(ctx context.Context, id string) (game.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(game.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockGamesServiceMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGamesService)(nil).Get), ctx, id)
}

// Reveal mocks base method.
func (m *MockGamesService) Reveal(ctx context.Context, id string, row, col int) (game.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reveal", ctx, id, row, col)
	ret0, _ := ret[0].(game.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reveal indicates an expected call of Reveal.
func (mr *MockGamesServiceMockRecorder) Reveal(ctx, id, row, col interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reveal", reflect.TypeOf((*MockGamesService)(nil).Reveal), ctx, id, row, col)
}
