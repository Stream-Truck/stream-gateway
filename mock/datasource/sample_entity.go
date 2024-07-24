// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/datasource/sample_entitiy/sample_entity.go
//
// Generated by this command:
//
//	mockgen -source=./internal/datasource/sample_entitiy/sample_entity.go -destination=./mock/datasource/sample_entity.go
//

// Package mock_sample_entitiy is a generated GoMock package.
package mock_sample_entitiy

import (
	context "context"
	reflect "reflect"

	"application/internal/v1/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockDataSource is a mock of DataSource interface.
type MockDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockDataSourceMockRecorder
}

// MockDataSourceMockRecorder is the mock recorder for MockDataSource.
type MockDataSourceMockRecorder struct {
	mock *MockDataSource
}

// NewMockDataSource creates a new mock instance.
func NewMockDataSource(ctrl *gomock.Controller) *MockDataSource {
	mock := &MockDataSource{ctrl: ctrl}
	mock.recorder = &MockDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSource) EXPECT() *MockDataSourceMockRecorder {
	return m.recorder
}

// Create mock base method.
func (m *MockDataSource) Create(ctx context.Context, sampleEntity *entity.SampleEntity) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, sampleEntity)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDataSourceMockRecorder) Create(ctx, sampleEntity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDataSource)(nil).Create), ctx, sampleEntity)
}

// Delete mock base method.
func (m *MockDataSource) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDataSourceMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDataSource)(nil).Delete), ctx, id)
}

// List mock base method.
func (m *MockDataSource) List(ctx context.Context) ([]*entity.SampleEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]*entity.SampleEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockDataSourceMockRecorder) List(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDataSource)(nil).List), ctx)
}

// Update mock base method.
func (m *MockDataSource) Update(ctx context.Context, sampleEntity *entity.SampleEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, sampleEntity)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockDataSourceMockRecorder) Update(ctx, sampleEntity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDataSource)(nil).Update), ctx, sampleEntity)
}
