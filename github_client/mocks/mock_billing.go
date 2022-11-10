package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-github/constants"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v45/github"
)

type MockBillingService struct {
	ctrl     *gomock.Controller
	recorder *MockBillingServiceMockRecorder
}

type MockBillingServiceMockRecorder struct {
	mock *MockBillingService
}

func NewMockBillingService(ctrl *gomock.Controller) *MockBillingService {
	mock := &MockBillingService{ctrl: ctrl}
	mock.recorder = &MockBillingServiceMockRecorder{mock}
	return mock
}

func (m *MockBillingService) EXPECT() *MockBillingServiceMockRecorder {
	return m.recorder
}

func (m *MockBillingService) GetActionsBillingOrg(arg0 context.Context, arg1 string) (*github.ActionBilling, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetActionsBillingOrg, arg0, arg1)
	ret0, _ := ret[0].(*github.ActionBilling)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockBillingServiceMockRecorder) GetActionsBillingOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetActionsBillingOrg, reflect.TypeOf((*MockBillingService)(nil).GetActionsBillingOrg), arg0, arg1)
}

func (m *MockBillingService) GetPackagesBillingOrg(arg0 context.Context, arg1 string) (*github.PackageBilling, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetPackagesBillingOrg, arg0, arg1)
	ret0, _ := ret[0].(*github.PackageBilling)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockBillingServiceMockRecorder) GetPackagesBillingOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetPackagesBillingOrg, reflect.TypeOf((*MockBillingService)(nil).GetPackagesBillingOrg), arg0, arg1)
}

func (m *MockBillingService) GetStorageBillingOrg(arg0 context.Context, arg1 string) (*github.StorageBilling, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetStorageBillingOrg, arg0, arg1)
	ret0, _ := ret[0].(*github.StorageBilling)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockBillingServiceMockRecorder) GetStorageBillingOrg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetStorageBillingOrg, reflect.TypeOf((*MockBillingService)(nil).GetStorageBillingOrg), arg0, arg1)
}
