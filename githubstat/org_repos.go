package githubstat

import (
	"fmt"

	"github.com/google/go-github/github"
)

func orgRepositoriesList() []github.Repository {
	opt := &github.RepositoryListByOrgOptions{
		Type: "all",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	proxyClient := &ProxyClient{}
	client := proxyClient.getClient()
	repos, _, err := client.Repositories.ListByOrg(config.OrgName, opt)

	if err != nil {
		fmt.Println(err)
	}

	return repos
}
