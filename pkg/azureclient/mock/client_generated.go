// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	dns "github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	gomock "github.com/golang/mock/gomock"
	azureclient "github.com/openshift/hive/pkg/azureclient"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ListResourceSKUs mocks base method
func (m *MockClient) ListResourceSKUs(ctx context.Context) (azureclient.ResourceSKUsPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceSKUs", ctx)
	ret0, _ := ret[0].(azureclient.ResourceSKUsPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceSKUs indicates an expected call of ListResourceSKUs
func (mr *MockClientMockRecorder) ListResourceSKUs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceSKUs", reflect.TypeOf((*MockClient)(nil).ListResourceSKUs), ctx)
}

// CreateOrUpdateZone mocks base method
func (m *MockClient) CreateOrUpdateZone(ctx context.Context, resourceGroupName, zone string) (dns.Zone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateZone", ctx, resourceGroupName, zone)
	ret0, _ := ret[0].(dns.Zone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateZone indicates an expected call of CreateOrUpdateZone
func (mr *MockClientMockRecorder) CreateOrUpdateZone(ctx, resourceGroupName, zone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateZone", reflect.TypeOf((*MockClient)(nil).CreateOrUpdateZone), ctx, resourceGroupName, zone)
}

// DeleteZone mocks base method
func (m *MockClient) DeleteZone(ctx context.Context, resourceGroupName, zone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteZone", ctx, resourceGroupName, zone)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteZone indicates an expected call of DeleteZone
func (mr *MockClientMockRecorder) DeleteZone(ctx, resourceGroupName, zone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteZone", reflect.TypeOf((*MockClient)(nil).DeleteZone), ctx, resourceGroupName, zone)
}

// GetZone mocks base method
func (m *MockClient) GetZone(ctx context.Context, resourceGroupName, zone string) (dns.Zone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZone", ctx, resourceGroupName, zone)
	ret0, _ := ret[0].(dns.Zone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetZone indicates an expected call of GetZone
func (mr *MockClientMockRecorder) GetZone(ctx, resourceGroupName, zone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZone", reflect.TypeOf((*MockClient)(nil).GetZone), ctx, resourceGroupName, zone)
}

// ListZones mocks base method
func (m *MockClient) ListZones(ctx context.Context, resourceGroupName string, top *int32) (dns.ZoneListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListZones", ctx, resourceGroupName, top)
	ret0, _ := ret[0].(dns.ZoneListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListZones indicates an expected call of ListZones
func (mr *MockClientMockRecorder) ListZones(ctx, resourceGroupName, top interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListZones", reflect.TypeOf((*MockClient)(nil).ListZones), ctx, resourceGroupName, top)
}

// ListRecordSetsByZone mocks base method
func (m *MockClient) ListRecordSetsByZone(ctx context.Context, resourceGroupName, zone string, top *int32) (dns.RecordSetListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecordSetsByZone", ctx, resourceGroupName, zone, top)
	ret0, _ := ret[0].(dns.RecordSetListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecordSetsByZone indicates an expected call of ListRecordSetsByZone
func (mr *MockClientMockRecorder) ListRecordSetsByZone(ctx, resourceGroupName, zone, top interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecordSetsByZone", reflect.TypeOf((*MockClient)(nil).ListRecordSetsByZone), ctx, resourceGroupName, zone, top)
}

// CreateOrUpdateRecordSet mocks base method
func (m *MockClient) CreateOrUpdateRecordSet(ctx context.Context, resourceGroupName, zone, recordSetName string, recordType dns.RecordType, recordSet dns.RecordSet) (dns.RecordSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateRecordSet", ctx, resourceGroupName, zone, recordSetName, recordType, recordSet)
	ret0, _ := ret[0].(dns.RecordSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateRecordSet indicates an expected call of CreateOrUpdateRecordSet
func (mr *MockClientMockRecorder) CreateOrUpdateRecordSet(ctx, resourceGroupName, zone, recordSetName, recordType, recordSet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateRecordSet", reflect.TypeOf((*MockClient)(nil).CreateOrUpdateRecordSet), ctx, resourceGroupName, zone, recordSetName, recordType, recordSet)
}

// DeleteRecordSet mocks base method
func (m *MockClient) DeleteRecordSet(ctx context.Context, resourceGroupName, zone, recordSetName string, recordType dns.RecordType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecordSet", ctx, resourceGroupName, zone, recordSetName, recordType)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecordSet indicates an expected call of DeleteRecordSet
func (mr *MockClientMockRecorder) DeleteRecordSet(ctx, resourceGroupName, zone, recordSetName, recordType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecordSet", reflect.TypeOf((*MockClient)(nil).DeleteRecordSet), ctx, resourceGroupName, zone, recordSetName, recordType)
}

// MockResourceSKUsPage is a mock of ResourceSKUsPage interface
type MockResourceSKUsPage struct {
	ctrl     *gomock.Controller
	recorder *MockResourceSKUsPageMockRecorder
}

// MockResourceSKUsPageMockRecorder is the mock recorder for MockResourceSKUsPage
type MockResourceSKUsPageMockRecorder struct {
	mock *MockResourceSKUsPage
}

// NewMockResourceSKUsPage creates a new mock instance
func NewMockResourceSKUsPage(ctrl *gomock.Controller) *MockResourceSKUsPage {
	mock := &MockResourceSKUsPage{ctrl: ctrl}
	mock.recorder = &MockResourceSKUsPageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResourceSKUsPage) EXPECT() *MockResourceSKUsPageMockRecorder {
	return m.recorder
}

// NextWithContext mocks base method
func (m *MockResourceSKUsPage) NextWithContext(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextWithContext", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// NextWithContext indicates an expected call of NextWithContext
func (mr *MockResourceSKUsPageMockRecorder) NextWithContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextWithContext", reflect.TypeOf((*MockResourceSKUsPage)(nil).NextWithContext), ctx)
}

// NotDone mocks base method
func (m *MockResourceSKUsPage) NotDone() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotDone")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NotDone indicates an expected call of NotDone
func (mr *MockResourceSKUsPageMockRecorder) NotDone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotDone", reflect.TypeOf((*MockResourceSKUsPage)(nil).NotDone))
}

// Values mocks base method
func (m *MockResourceSKUsPage) Values() []compute.ResourceSku {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Values")
	ret0, _ := ret[0].([]compute.ResourceSku)
	return ret0
}

// Values indicates an expected call of Values
func (mr *MockResourceSKUsPageMockRecorder) Values() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockResourceSKUsPage)(nil).Values))
}
