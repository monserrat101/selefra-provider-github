package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-github/constants"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v45/github"
)

type MockTeamsService struct {
	ctrl     *gomock.Controller
	recorder *MockTeamsServiceMockRecorder
}

type MockTeamsServiceMockRecorder struct {
	mock *MockTeamsService
}

func NewMockTeamsService(ctrl *gomock.Controller) *MockTeamsService {
	mock := &MockTeamsService{ctrl: ctrl}
	mock.recorder = &MockTeamsServiceMockRecorder{mock}
	return mock
}

func (m *MockTeamsService) EXPECT() *MockTeamsServiceMockRecorder {
	return m.recorder
}

func (m *MockTeamsService) GetTeamMembershipBySlug(arg0 context.Context, arg1, arg2, arg3 string) (*github.Membership, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetTeamMembershipBySlug, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Membership)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockTeamsServiceMockRecorder) GetTeamMembershipBySlug(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetTeamMembershipBySlug, reflect.TypeOf((*MockTeamsService)(nil).GetTeamMembershipBySlug), arg0, arg1, arg2, arg3)
}

func (m *MockTeamsService) ListExternalGroups(arg0 context.Context, arg1 string, arg2 *github.ListExternalGroupsOptions) (*github.ExternalGroupList, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListExternalGroups, arg0, arg1, arg2)
	ret0, _ := ret[0].(*github.ExternalGroupList)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockTeamsServiceMockRecorder) ListExternalGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListExternalGroups, reflect.TypeOf((*MockTeamsService)(nil).ListExternalGroups), arg0, arg1, arg2)
}

func (m *MockTeamsService) ListTeamMembersByID(arg0 context.Context, arg1, arg2 int64, arg3 *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListTeamMembersByID, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockTeamsServiceMockRecorder) ListTeamMembersByID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTeamMembersByID, reflect.TypeOf((*MockTeamsService)(nil).ListTeamMembersByID), arg0, arg1, arg2, arg3)
}

func (m *MockTeamsService) ListTeamReposByID(arg0 context.Context, arg1, arg2 int64, arg3 *github.ListOptions) ([]*github.Repository, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListTeamReposByID, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*github.Repository)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockTeamsServiceMockRecorder) ListTeamReposByID(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTeamReposByID, reflect.TypeOf((*MockTeamsService)(nil).ListTeamReposByID), arg0, arg1, arg2, arg3)
}

func (m *MockTeamsService) ListTeams(arg0 context.Context, arg1 string, arg2 *github.ListOptions) ([]*github.Team, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListTeams, arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Team)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockTeamsServiceMockRecorder) ListTeams(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTeams, reflect.TypeOf((*MockTeamsService)(nil).ListTeams), arg0, arg1, arg2)
}
