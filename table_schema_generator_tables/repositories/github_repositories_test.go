package repositories

import (
	"github.com/selefra/selefra-provider-github/constants"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v45/github"

	"github.com/selefra/selefra-provider-github/faker"
	"github.com/selefra/selefra-provider-github/github_client"

	"github.com/selefra/selefra-provider-github/github_client/mocks"

	"github.com/selefra/selefra-provider-github/table_schema_generator"
)

func buildRepositiories(t *testing.T, ctrl *gomock.Controller) github_client.GithubServices {
	mock := mocks.NewMockRepositoriesService(ctrl)

	var cs github.Repository

	if err := faker.FakeObject(&cs); err != nil {

		t.Fatal(err)

	}

	someId := int64(5555555)

	cs.Parent = &github.Repository{ID: &someId}
	cs.TemplateRepository = &github.Repository{ID: &someId}

	cs.Source = &github.Repository{ID: &someId}

	mock.EXPECT().ListByOrg(gomock.Any(), constants.Testorg, gomock.Any()).AnyTimes().Return(

		[]*github.Repository{&cs}, &github.Response{}, nil)

	return github_client.GithubServices{Repositories: mock}
}

func TestRepos(t *testing.T) {

	github_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGithubRepositoriesGenerator{}), buildRepositiories, github_client.TestOptions{})

}
