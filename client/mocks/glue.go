// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: GlueClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	glue "github.com/aws/aws-sdk-go-v2/service/glue"
	gomock "github.com/golang/mock/gomock"
)

// MockGlueClient is a mock of GlueClient interface.
type MockGlueClient struct {
	ctrl     *gomock.Controller
	recorder *MockGlueClientMockRecorder
}

// MockGlueClientMockRecorder is the mock recorder for MockGlueClient.
type MockGlueClientMockRecorder struct {
	mock *MockGlueClient
}

// NewMockGlueClient creates a new mock instance.
func NewMockGlueClient(ctrl *gomock.Controller) *MockGlueClient {
	mock := &MockGlueClient{ctrl: ctrl}
	mock.recorder = &MockGlueClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGlueClient) EXPECT() *MockGlueClientMockRecorder {
	return m.recorder
}

// GetTags mocks base method.
func (m *MockGlueClient) GetTags(arg0 context.Context, arg1 *glue.GetTagsInput, arg2 ...func(*glue.Options)) (*glue.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTags", varargs...)
	ret0, _ := ret[0].(*glue.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags.
func (mr *MockGlueClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockGlueClient)(nil).GetTags), varargs...)
}

// GetWorkflow mocks base method.
func (m *MockGlueClient) GetWorkflow(arg0 context.Context, arg1 *glue.GetWorkflowInput, arg2 ...func(*glue.Options)) (*glue.GetWorkflowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflow", varargs...)
	ret0, _ := ret[0].(*glue.GetWorkflowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkflow indicates an expected call of GetWorkflow.
func (mr *MockGlueClientMockRecorder) GetWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflow", reflect.TypeOf((*MockGlueClient)(nil).GetWorkflow), varargs...)
}

// ListWorkflows mocks base method.
func (m *MockGlueClient) ListWorkflows(arg0 context.Context, arg1 *glue.ListWorkflowsInput, arg2 ...func(*glue.Options)) (*glue.ListWorkflowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWorkflows", varargs...)
	ret0, _ := ret[0].(*glue.ListWorkflowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkflows indicates an expected call of ListWorkflows.
func (mr *MockGlueClientMockRecorder) ListWorkflows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflows", reflect.TypeOf((*MockGlueClient)(nil).ListWorkflows), varargs...)
}