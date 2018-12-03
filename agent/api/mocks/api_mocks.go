// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/api (interfaces: ECSSDK,ECSSubmitStateSDK,ECSClient)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	api "github.com/aws/amazon-ecs-agent/agent/api"
	ecs "github.com/aws/amazon-ecs-agent/agent/ecs_client/model/ecs"
	gomock "github.com/golang/mock/gomock"
)

// MockECSSDK is a mock of ECSSDK interface
type MockECSSDK struct {
	ctrl     *gomock.Controller
	recorder *MockECSSDKMockRecorder
}

// MockECSSDKMockRecorder is the mock recorder for MockECSSDK
type MockECSSDKMockRecorder struct {
	mock *MockECSSDK
}

// NewMockECSSDK creates a new mock instance
func NewMockECSSDK(ctrl *gomock.Controller) *MockECSSDK {
	mock := &MockECSSDK{ctrl: ctrl}
	mock.recorder = &MockECSSDKMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECSSDK) EXPECT() *MockECSSDKMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockECSSDK) CreateCluster(arg0 *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error) {
	ret := m.ctrl.Call(m, "CreateCluster", arg0)
	ret0, _ := ret[0].(*ecs.CreateClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockECSSDKMockRecorder) CreateCluster(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockECSSDK)(nil).CreateCluster), arg0)
}

// DiscoverPollEndpoint mocks base method
func (m *MockECSSDK) DiscoverPollEndpoint(arg0 *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error) {
	ret := m.ctrl.Call(m, "DiscoverPollEndpoint", arg0)
	ret0, _ := ret[0].(*ecs.DiscoverPollEndpointOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverPollEndpoint indicates an expected call of DiscoverPollEndpoint
func (mr *MockECSSDKMockRecorder) DiscoverPollEndpoint(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverPollEndpoint", reflect.TypeOf((*MockECSSDK)(nil).DiscoverPollEndpoint), arg0)
}

// RegisterContainerInstance mocks base method
func (m *MockECSSDK) RegisterContainerInstance(arg0 *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error) {
	ret := m.ctrl.Call(m, "RegisterContainerInstance", arg0)
	ret0, _ := ret[0].(*ecs.RegisterContainerInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterContainerInstance indicates an expected call of RegisterContainerInstance
func (mr *MockECSSDKMockRecorder) RegisterContainerInstance(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterContainerInstance", reflect.TypeOf((*MockECSSDK)(nil).RegisterContainerInstance), arg0)
}

// MockECSSubmitStateSDK is a mock of ECSSubmitStateSDK interface
type MockECSSubmitStateSDK struct {
	ctrl     *gomock.Controller
	recorder *MockECSSubmitStateSDKMockRecorder
}

// MockECSSubmitStateSDKMockRecorder is the mock recorder for MockECSSubmitStateSDK
type MockECSSubmitStateSDKMockRecorder struct {
	mock *MockECSSubmitStateSDK
}

// NewMockECSSubmitStateSDK creates a new mock instance
func NewMockECSSubmitStateSDK(ctrl *gomock.Controller) *MockECSSubmitStateSDK {
	mock := &MockECSSubmitStateSDK{ctrl: ctrl}
	mock.recorder = &MockECSSubmitStateSDKMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECSSubmitStateSDK) EXPECT() *MockECSSubmitStateSDKMockRecorder {
	return m.recorder
}

// SubmitContainerStateChange mocks base method
func (m *MockECSSubmitStateSDK) SubmitContainerStateChange(arg0 *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error) {
	ret := m.ctrl.Call(m, "SubmitContainerStateChange", arg0)
	ret0, _ := ret[0].(*ecs.SubmitContainerStateChangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitContainerStateChange indicates an expected call of SubmitContainerStateChange
func (mr *MockECSSubmitStateSDKMockRecorder) SubmitContainerStateChange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitContainerStateChange", reflect.TypeOf((*MockECSSubmitStateSDK)(nil).SubmitContainerStateChange), arg0)
}

// SubmitTaskStateChange mocks base method
func (m *MockECSSubmitStateSDK) SubmitTaskStateChange(arg0 *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error) {
	ret := m.ctrl.Call(m, "SubmitTaskStateChange", arg0)
	ret0, _ := ret[0].(*ecs.SubmitTaskStateChangeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTaskStateChange indicates an expected call of SubmitTaskStateChange
func (mr *MockECSSubmitStateSDKMockRecorder) SubmitTaskStateChange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTaskStateChange", reflect.TypeOf((*MockECSSubmitStateSDK)(nil).SubmitTaskStateChange), arg0)
}

// MockECSClient is a mock of ECSClient interface
type MockECSClient struct {
	ctrl     *gomock.Controller
	recorder *MockECSClientMockRecorder
}

// MockECSClientMockRecorder is the mock recorder for MockECSClient
type MockECSClientMockRecorder struct {
	mock *MockECSClient
}

// NewMockECSClient creates a new mock instance
func NewMockECSClient(ctrl *gomock.Controller) *MockECSClient {
	mock := &MockECSClient{ctrl: ctrl}
	mock.recorder = &MockECSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECSClient) EXPECT() *MockECSClientMockRecorder {
	return m.recorder
}

// DiscoverPollEndpoint mocks base method
func (m *MockECSClient) DiscoverPollEndpoint(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "DiscoverPollEndpoint", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverPollEndpoint indicates an expected call of DiscoverPollEndpoint
func (mr *MockECSClientMockRecorder) DiscoverPollEndpoint(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverPollEndpoint", reflect.TypeOf((*MockECSClient)(nil).DiscoverPollEndpoint), arg0)
}

// DiscoverTelemetryEndpoint mocks base method
func (m *MockECSClient) DiscoverTelemetryEndpoint(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "DiscoverTelemetryEndpoint", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverTelemetryEndpoint indicates an expected call of DiscoverTelemetryEndpoint
func (mr *MockECSClientMockRecorder) DiscoverTelemetryEndpoint(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverTelemetryEndpoint", reflect.TypeOf((*MockECSClient)(nil).DiscoverTelemetryEndpoint), arg0)
}

// RegisterContainerInstance mocks base method
func (m *MockECSClient) RegisterContainerInstance(arg0 string, arg1 []*ecs.Attribute, arg2 []*ecs.Tag, arg3 string) (string, string, error) {
	ret := m.ctrl.Call(m, "RegisterContainerInstance", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterContainerInstance indicates an expected call of RegisterContainerInstance
func (mr *MockECSClientMockRecorder) RegisterContainerInstance(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterContainerInstance", reflect.TypeOf((*MockECSClient)(nil).RegisterContainerInstance), arg0, arg1, arg2, arg3)
}

// SubmitContainerStateChange mocks base method
func (m *MockECSClient) SubmitContainerStateChange(arg0 api.ContainerStateChange) error {
	ret := m.ctrl.Call(m, "SubmitContainerStateChange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitContainerStateChange indicates an expected call of SubmitContainerStateChange
func (mr *MockECSClientMockRecorder) SubmitContainerStateChange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitContainerStateChange", reflect.TypeOf((*MockECSClient)(nil).SubmitContainerStateChange), arg0)
}

// SubmitTaskStateChange mocks base method
func (m *MockECSClient) SubmitTaskStateChange(arg0 api.TaskStateChange) error {
	ret := m.ctrl.Call(m, "SubmitTaskStateChange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitTaskStateChange indicates an expected call of SubmitTaskStateChange
func (mr *MockECSClientMockRecorder) SubmitTaskStateChange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTaskStateChange", reflect.TypeOf((*MockECSClient)(nil).SubmitTaskStateChange), arg0)
}
