package organizations

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v45/github"

	"github.com/selefra/selefra-provider-github/faker"
	"github.com/selefra/selefra-provider-github/github_client"

	"github.com/selefra/selefra-provider-github/github_client/mocks"

	"github.com/selefra/selefra-provider-github/table_schema_generator"
)

func buildOrganizations(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {

	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs *github.Organization

	if err := faker.FakeObject(&cs); err != nil {

		t.Fatal(err)
	}

	mock.EXPECT().Get(gomock.Any(), gomock.Any()).AnyTimes().Return(cs, &github.Response{}, nil)

	var u github.User
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)

	}
	mock.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		[]*github.User{&u}, &github.Response{}, nil)

	var m github.Membership

	if err := faker.FakeObject(&m); err != nil {

		t.Fatal(err)

	}
	mock.EXPECT().GetOrgMembership(gomock.Any(), *u.Login, gomock.Any()).AnyTimes().Return(
		&m, &github.Response{}, nil)

	return github_client.GithubServices{Organizations: mock}

}

func TestOrganizations(t *testing.T) {
	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubOrganizationsGenerator{}), buildOrganizations, github_client.TestOptions{})
}
