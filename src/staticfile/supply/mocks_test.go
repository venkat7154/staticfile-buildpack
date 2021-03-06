// Automatically generated by MockGen. DO NOT EDIT!
// Source: supply.go

package supply_test

import (
	libbuildpack "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Manifest interface
type MockManifest struct {
	ctrl     *gomock.Controller
	recorder *_MockManifestRecorder
}

// Recorder for MockManifest (not exported)
type _MockManifestRecorder struct {
	mock *MockManifest
}

func NewMockManifest(ctrl *gomock.Controller) *MockManifest {
	mock := &MockManifest{ctrl: ctrl}
	mock.recorder = &_MockManifestRecorder{mock}
	return mock
}

func (_m *MockManifest) EXPECT() *_MockManifestRecorder {
	return _m.recorder
}

func (_m *MockManifest) DefaultVersion(_param0 string) (libbuildpack.Dependency, error) {
	ret := _m.ctrl.Call(_m, "DefaultVersion", _param0)
	ret0, _ := ret[0].(libbuildpack.Dependency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManifestRecorder) DefaultVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DefaultVersion", arg0)
}

func (_m *MockManifest) InstallDependency(_param0 libbuildpack.Dependency, _param1 string) error {
	ret := _m.ctrl.Call(_m, "InstallDependency", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManifestRecorder) InstallDependency(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InstallDependency", arg0, arg1)
}

// Mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *_MockStagerRecorder
}

// Recorder for MockStager (not exported)
type _MockStagerRecorder struct {
	mock *MockStager
}

func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &_MockStagerRecorder{mock}
	return mock
}

func (_m *MockStager) EXPECT() *_MockStagerRecorder {
	return _m.recorder
}

func (_m *MockStager) AddBinDependencyLink(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "AddBinDependencyLink", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStagerRecorder) AddBinDependencyLink(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddBinDependencyLink", arg0, arg1)
}

func (_m *MockStager) DepDir() string {
	ret := _m.ctrl.Call(_m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockStagerRecorder) DepDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DepDir")
}
