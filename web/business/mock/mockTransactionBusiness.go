// Code generated by MockGen. DO NOT EDIT.
// Source: web/business/transactionBusiness.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/TulioGuaraldoB/q2-payer-challenge/web/model"
	gomock "github.com/golang/mock/gomock"
)

// MockITransactionBusiness is a mock of ITransactionBusiness interface.
type MockITransactionBusiness struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionBusinessMockRecorder
}

// MockITransactionBusinessMockRecorder is the mock recorder for MockITransactionBusiness.
type MockITransactionBusinessMockRecorder struct {
	mock *MockITransactionBusiness
}

// NewMockITransactionBusiness creates a new mock instance.
func NewMockITransactionBusiness(ctrl *gomock.Controller) *MockITransactionBusiness {
	mock := &MockITransactionBusiness{ctrl: ctrl}
	mock.recorder = &MockITransactionBusinessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionBusiness) EXPECT() *MockITransactionBusinessMockRecorder {
	return m.recorder
}

// GetTransactionById mocks base method.
func (m *MockITransactionBusiness) GetTransactionById(transactionId uint) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionById", transactionId)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionById indicates an expected call of GetTransactionById.
func (mr *MockITransactionBusinessMockRecorder) GetTransactionById(transactionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionById", reflect.TypeOf((*MockITransactionBusiness)(nil).GetTransactionById), transactionId)
}