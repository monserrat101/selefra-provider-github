package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-github/constants"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v45/github"
)

type MockIssuesService struct {
	ctrl     *gomock.Controller
	recorder *MockIssuesServiceMockRecorder
}

type MockIssuesServiceMockRecorder struct {
	mock *MockIssuesService
}

func NewMockIssuesService(ctrl *gomock.Controller) *MockIssuesService {
	mock := &MockIssuesService{ctrl: ctrl}
	mock.recorder = &MockIssuesServiceMockRecorder{mock}
	return mock
}

func (m *MockIssuesService) EXPECT() *MockIssuesServiceMockRecorder {
	return m.recorder
}

func (m *MockIssuesService) ListByOrg(arg0 context.Context, arg1 string, arg2 *github.IssueListOptions) ([]*github.Issue, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListByOrg, arg0, arg1, arg2)
	ret0, _ := ret[0].([]*github.Issue)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockIssuesServiceMockRecorder) ListByOrg(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListByOrg, reflect.TypeOf((*MockIssuesService)(nil).ListByOrg), arg0, arg1, arg2)
}
