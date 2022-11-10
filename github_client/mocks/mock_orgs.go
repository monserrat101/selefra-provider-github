package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-github/constants"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v45/github"
)

type MockOrganizationsService struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsServiceMockRecorder
}

type MockOrganizationsServiceMockRecorder struct {
	mock *MockOrganizationsService
}

func NewMockOrganizationsService(ctrl *gomock.Controller) *MockOrganizationsService {
	mock := &MockOrganizationsService{ctrl: ctrl}
	mock.recorder = &MockOrganizationsServiceMockRecorder{mock}
	return mock
}

func (m *MockOrganizationsService) EXPECT() *MockOrganizationsServiceMockRecorder {
	return m.recorder
}

func (m *MockOrganizationsService) Get(arg0 context.Context, arg1 string) (*github.Organization, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0, arg1)
	ret0, _ := ret[0].(*github.Organization)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockOrganizationsServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockOrganizationsService)(nil).Get), arg0, arg1)
}

func (m *MockOrganizationsService) GetOrgMembership(arg0 context.Context, arg1, arg2 string) (*github.Membership, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetOrgMembership, arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.Membership)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockOrganizationsServiceMockRecorder) GetOrgMembership(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetOrgMembership, reflect.TypeOf((*MockOrganizationsService)(nil).GetOrgMembership), arg0, arg1, arg2)
}

func (m *MockOrganizationsService) ListHookDeliveries(arg0 context.Context, arg1 string, arg2 int64, arg3 *github.ListCursorOptions) ([]*github.HookDelivery, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListHookDeliveries, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.HookDelivery)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockOrganizationsServiceMockRecorder) ListHookDeliveries(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHookDeliveries, reflect.TypeOf((*MockOrganizationsService)(nil).ListHookDeliveries), arg0, arg1, arg2, arg3)
}

func (m *MockOrganizationsService) ListHooks(arg0 context.Context, arg1 string, arg2 *github.ListOptions) ([]*github.Hook, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListHooks, arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Hook)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockOrganizationsServiceMockRecorder) ListHooks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListHooks, reflect.TypeOf((*MockOrganizationsService)(nil).ListHooks), arg0, arg1, arg2)
}

func (m *MockOrganizationsService) ListInstallations(arg0 context.Context, arg1 string, arg2 *github.ListOptions) (*github.OrganizationInstallations, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListInstallations, arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.OrganizationInstallations)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockOrganizationsServiceMockRecorder) ListInstallations(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListInstallations, reflect.TypeOf((*MockOrganizationsService)(nil).ListInstallations), arg0, arg1, arg2)
}

func (m *MockOrganizationsService) ListMembers(arg0 context.Context, arg1 string, arg2 *github.ListMembersOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListMembers, arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockOrganizationsServiceMockRecorder) ListMembers(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMembers, reflect.TypeOf((*MockOrganizationsService)(nil).ListMembers), arg0, arg1, arg2)
}
