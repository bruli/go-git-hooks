package main

import (
	"github.com/bruli/go-git-hooks/configuration"
	"github.com/bruli/go-git-hooks/installerActions"
	"log"
)

func main() {
	preCommitConf := configuration.PreCommitConfiguration{}
	err := installerActions.PreCommitAction(&preCommitConf)
	if err != nil {
		log.Fatal(err)
	}

	conf := configuration.Configuration{PreCommit: &preCommitConf}
	if err = configuration.WriteConfiguration(&conf); err != nil {
		log.Fatal(err)
	}
}
