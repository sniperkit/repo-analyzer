package detectors

import (
	"encoding/json"
	"log"

	"github.com/pvaass/repo-analyzer/pkg/repository"
)

type ComposerDependencyDetector struct {
	composer ComposerDependencyMap
}

type ComposerDependencyMap struct {
	Name    string
	Require map[string]string
}

func (f ComposerDependencyDetector) Supports(rule Rule) bool {
	return rule.Strategy == "composer#d" || rule.Strategy == "composer#f"
}

func (f *ComposerDependencyDetector) Init(repo repository.Repository) {
	file := findFile(
		repo,
		[]string{
			"composer.json",
			"app/composer.json",
		},
	)
	if len(file) <= 0 {
		return
	}

	err := json.Unmarshal(file, &f.composer)
	if err != nil {
		log.Panic("Invalid Json Decode", err)
	}
}

func (f ComposerDependencyDetector) Detect(repo repository.Repository, resultChannel chan Result, rule Rule) {
	result := Result{Identifier: rule.Name}

	var ok bool

	switch rule.Strategy {
	case "composer#d":
		_, ok = f.composer.Require[rule.Arguments[0]]
	case "composer#f":
		if rule.Arguments[0] == "name" {
			ok = f.composer.Name == rule.Arguments[1]
		}
	}

	if ok {
		result.Score = 100
	}
	resultChannel <- result
}

func init() {
	register(&ComposerDependencyDetector{})
}
