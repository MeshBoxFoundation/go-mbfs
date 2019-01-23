// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lucas-clemente/quic-go (interfaces: QuicAEAD)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "mbfs/go-mbfs/gx/QmU44KWVkSHno7sNDTeUcL4FBgxgoidkFuTUyTXWJPXXFJ/quic-go/internal/protocol"
)

// MockQuicAEAD is a mock of QuicAEAD interface
type MockQuicAEAD struct {
	ctrl     *gomock.Controller
	recorder *MockQuicAEADMockRecorder
}

// MockQuicAEADMockRecorder is the mock recorder for MockQuicAEAD
type MockQuicAEADMockRecorder struct {
	mock *MockQuicAEAD
}

// NewMockQuicAEAD creates a new mock instance
func NewMockQuicAEAD(ctrl *gomock.Controller) *MockQuicAEAD {
	mock := &MockQuicAEAD{ctrl: ctrl}
	mock.recorder = &MockQuicAEADMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuicAEAD) EXPECT() *MockQuicAEADMockRecorder {
	return m.recorder
}

// Open1RTT mocks base method
func (m *MockQuicAEAD) Open1RTT(arg0, arg1 []byte, arg2 protocol.PacketNumber, arg3 []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "Open1RTT", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open1RTT indicates an expected call of Open1RTT
func (mr *MockQuicAEADMockRecorder) Open1RTT(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open1RTT", reflect.TypeOf((*MockQuicAEAD)(nil).Open1RTT), arg0, arg1, arg2, arg3)
}

// OpenHandshake mocks base method
func (m *MockQuicAEAD) OpenHandshake(arg0, arg1 []byte, arg2 protocol.PacketNumber, arg3 []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "OpenHandshake", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenHandshake indicates an expected call of OpenHandshake
func (mr *MockQuicAEADMockRecorder) OpenHandshake(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenHandshake", reflect.TypeOf((*MockQuicAEAD)(nil).OpenHandshake), arg0, arg1, arg2, arg3)
}
