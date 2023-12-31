// Code generated by MockGen. DO NOT EDIT.
// Source: x/checkers/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBankEscrowKeeper is a mock of BankEscrowKeeper interface.
type MockBankEscrowKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankEscrowKeeperMockRecorder
}

// MockBankEscrowKeeperMockRecorder is the mock recorder for MockBankEscrowKeeper.
type MockBankEscrowKeeperMockRecorder struct {
	mock *MockBankEscrowKeeper
}

// NewMockBankEscrowKeeper creates a new mock instance.
func NewMockBankEscrowKeeper(ctrl *gomock.Controller) *MockBankEscrowKeeper {
	mock := &MockBankEscrowKeeper{ctrl: ctrl}
	mock.recorder = &MockBankEscrowKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankEscrowKeeper) EXPECT() *MockBankEscrowKeeperMockRecorder {
	return m.recorder
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankEscrowKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankEscrowKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankEscrowKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankEscrowKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankEscrowKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankEscrowKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}
