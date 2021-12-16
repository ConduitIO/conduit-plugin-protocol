// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit-plugin/cpluginv1 (interfaces: DestinationPlugin)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	cpluginv1 "github.com/conduitio/conduit-plugin/cpluginv1"
	gomock "github.com/golang/mock/gomock"
)

// DestinationPlugin is a mock of DestinationPlugin interface.
type DestinationPlugin struct {
	ctrl     *gomock.Controller
	recorder *DestinationPluginMockRecorder
}

// DestinationPluginMockRecorder is the mock recorder for DestinationPlugin.
type DestinationPluginMockRecorder struct {
	mock *DestinationPlugin
}

// NewDestinationPlugin creates a new mock instance.
func NewDestinationPlugin(ctrl *gomock.Controller) *DestinationPlugin {
	mock := &DestinationPlugin{ctrl: ctrl}
	mock.recorder = &DestinationPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *DestinationPlugin) EXPECT() *DestinationPluginMockRecorder {
	return m.recorder
}

// Configure mocks base method.
func (m *DestinationPlugin) Configure(arg0 context.Context, arg1 cpluginv1.DestinationConfigureRequest) (cpluginv1.DestinationConfigureResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(cpluginv1.DestinationConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure.
func (mr *DestinationPluginMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*DestinationPlugin)(nil).Configure), arg0, arg1)
}

// Run mocks base method.
func (m *DestinationPlugin) Run(arg0 context.Context, arg1 cpluginv1.DestinationRunStream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *DestinationPluginMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*DestinationPlugin)(nil).Run), arg0, arg1)
}

// Start mocks base method.
func (m *DestinationPlugin) Start(arg0 context.Context, arg1 cpluginv1.DestinationStartRequest) (cpluginv1.DestinationStartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(cpluginv1.DestinationStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *DestinationPluginMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*DestinationPlugin)(nil).Start), arg0, arg1)
}

// Stop mocks base method.
func (m *DestinationPlugin) Stop(arg0 context.Context, arg1 cpluginv1.DestinationStopRequest) (cpluginv1.DestinationStopResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(cpluginv1.DestinationStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *DestinationPluginMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*DestinationPlugin)(nil).Stop), arg0, arg1)
}

// Teardown mocks base method.
func (m *DestinationPlugin) Teardown(arg0 context.Context, arg1 cpluginv1.DestinationTeardownRequest) (cpluginv1.DestinationTeardownResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teardown", arg0, arg1)
	ret0, _ := ret[0].(cpluginv1.DestinationTeardownResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Teardown indicates an expected call of Teardown.
func (mr *DestinationPluginMockRecorder) Teardown(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teardown", reflect.TypeOf((*DestinationPlugin)(nil).Teardown), arg0, arg1)
}
