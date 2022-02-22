// Code generated by MockGen. DO NOT EDIT.
// Source: registry.go

// Package storage is a generated GoMock package.
package storage

import (
	reflect "reflect"

	v1alpha1 "github.com/HwameiStor/local-storage/pkg/apis/uds/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockLocalRegistry is a mock of LocalRegistry interface.
type MockLocalRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockLocalRegistryMockRecorder
}

// MockLocalRegistryMockRecorder is the mock recorder for MockLocalRegistry.
type MockLocalRegistryMockRecorder struct {
	mock *MockLocalRegistry
}

// NewMockLocalRegistry creates a new mock instance.
func NewMockLocalRegistry(ctrl *gomock.Controller) *MockLocalRegistry {
	mock := &MockLocalRegistry{ctrl: ctrl}
	mock.recorder = &MockLocalRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalRegistry) EXPECT() *MockLocalRegistryMockRecorder {
	return m.recorder
}

// Disks mocks base method.
func (m *MockLocalRegistry) Disks() map[string]*v1alpha1.LocalDisk {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disks")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalDisk)
	return ret0
}

// Disks indicates an expected call of Disks.
func (mr *MockLocalRegistryMockRecorder) Disks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disks", reflect.TypeOf((*MockLocalRegistry)(nil).Disks))
}

// HasVolumeReplica mocks base method.
func (m *MockLocalRegistry) HasVolumeReplica(replica *v1alpha1.LocalVolumeReplica) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasVolumeReplica", replica)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasVolumeReplica indicates an expected call of HasVolumeReplica.
func (mr *MockLocalRegistryMockRecorder) HasVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasVolumeReplica", reflect.TypeOf((*MockLocalRegistry)(nil).HasVolumeReplica), replica)
}

// Init mocks base method.
func (m *MockLocalRegistry) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init.
func (mr *MockLocalRegistryMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockLocalRegistry)(nil).Init))
}

// Pools mocks base method.
func (m *MockLocalRegistry) Pools() map[string]*v1alpha1.LocalPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pools")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalPool)
	return ret0
}

// Pools indicates an expected call of Pools.
func (mr *MockLocalRegistryMockRecorder) Pools() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pools", reflect.TypeOf((*MockLocalRegistry)(nil).Pools))
}

// SyncResourcesToNodeCRD mocks base method.
func (m *MockLocalRegistry) SyncResourcesToNodeCRD(localDisks map[string]*v1alpha1.LocalDisk) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncResourcesToNodeCRD", localDisks)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncResourcesToNodeCRD indicates an expected call of SyncResourcesToNodeCRD.
func (mr *MockLocalRegistryMockRecorder) SyncResourcesToNodeCRD(localDisks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncResourcesToNodeCRD", reflect.TypeOf((*MockLocalRegistry)(nil).SyncResourcesToNodeCRD), localDisks)
}

// UpdateNodeForVolumeReplica mocks base method.
func (m *MockLocalRegistry) UpdateNodeForVolumeReplica(replica *v1alpha1.LocalVolumeReplica) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateNodeForVolumeReplica", replica)
}

// UpdateNodeForVolumeReplica indicates an expected call of UpdateNodeForVolumeReplica.
func (mr *MockLocalRegistryMockRecorder) UpdateNodeForVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNodeForVolumeReplica", reflect.TypeOf((*MockLocalRegistry)(nil).UpdateNodeForVolumeReplica), replica)
}

// VolumeReplicas mocks base method.
func (m *MockLocalRegistry) VolumeReplicas() map[string]*v1alpha1.LocalVolumeReplica {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeReplicas")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalVolumeReplica)
	return ret0
}

// VolumeReplicas indicates an expected call of VolumeReplicas.
func (mr *MockLocalRegistryMockRecorder) VolumeReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeReplicas", reflect.TypeOf((*MockLocalRegistry)(nil).VolumeReplicas))
}