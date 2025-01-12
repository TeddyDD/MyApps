package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
	"icikowski.pl/myapps/common"
	"icikowski.pl/myapps/types"
)

var repos = types.Repositories{
	{
		Name:        "default",
		Description: "Default empty repository",
		Maintainer:  "MyApps",
	},
}

func init() {
	files, _ := filepath.Glob(common.PathRepositories + "/*.yaml")

	if len(files) > 0 {
		repos = types.Repositories{}
	}

	for _, file := range files {
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		var repo *types.Repository
		if err = yaml.Unmarshal(contents, &repo); err != nil {
			panic(err)
		}

		repos = append(repos, *repo)
	}
}

func GetRepositories() types.Repositories {
	return repos
}
