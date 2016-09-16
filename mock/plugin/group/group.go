// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/libmachete/plugin/group (interfaces: Scaled)

package group

import (
	instance "github.com/docker/libmachete/spi/instance"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Scaled interface
type MockScaled struct {
	ctrl     *gomock.Controller
	recorder *_MockScaledRecorder
}

// Recorder for MockScaled (not exported)
type _MockScaledRecorder struct {
	mock *MockScaled
}

func NewMockScaled(ctrl *gomock.Controller) *MockScaled {
	mock := &MockScaled{ctrl: ctrl}
	mock.recorder = &_MockScaledRecorder{mock}
	return mock
}

func (_m *MockScaled) EXPECT() *_MockScaledRecorder {
	return _m.recorder
}

func (_m *MockScaled) CreateOne(_param0 *string) {
	_m.ctrl.Call(_m, "CreateOne", _param0)
}

func (_mr *_MockScaledRecorder) CreateOne(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateOne", arg0)
}

func (_m *MockScaled) Destroy(_param0 instance.ID) {
	_m.ctrl.Call(_m, "Destroy", _param0)
}

func (_mr *_MockScaledRecorder) Destroy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Destroy", arg0)
}

func (_m *MockScaled) List() ([]instance.Description, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]instance.Description)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockScaledRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}