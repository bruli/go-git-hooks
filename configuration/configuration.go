package configuration

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Tool struct {
	Title string `yaml:"title"`
	Command string `yaml:"command"`
}

type Configuration struct {
	PreCommit *PreCommitConfiguration `yaml:"pre-commit"`
}

type PreCommitConfiguration struct {
	Tools []*Tool
}

func (p *PreCommitConfiguration) AddTool(t *Tool)  {
	p.Tools = append(p.Tools, t)
}

func WriteConfiguration(c *Configuration) error {
	d, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile("git-hooks.yaml", d, 0644); err != nil {
		return err
	}

	return nil
}
