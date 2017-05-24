package detectors

import (
	"encoding/json"
	"log"

	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type Laravel struct {
	composer struct {
		Require struct {
			Laravel string `json:"laravel/framework"`
		} `json:"require"`
	}
}

func (Laravel) Identifier() string {
	return "laravel"
}

func (f Laravel) Detect(repo repository.Repository) int {
	hasComposer := Composer{}.Detect(repo) >= 100
	if !hasComposer {
		return 0
	}

	f.getComposer(repo)

	if f.composer.Require.Laravel != "" {
		return 100
	}

	return 0
}

func (f *Laravel) getComposer(repo repository.Repository) {
	file := repo.File("composer.json")

	err := json.Unmarshal([]byte(file), &f.composer)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}

}