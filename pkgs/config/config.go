package config

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"
)

type Config struct {
	Name    string
	Configs map[string]interface{} `hcl:"config"`
	Addons  []string               `hcl:"addons"`
}

func Load(s string) (Config, error) {
	c := Config{}
	if err := hcl.Decode(&c, s); err != nil {
		return c, err
	}
	return c, nil
}

func LoadFile(path string) (Config, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, errors.Wrap(err, "couldn't load minishift config")
	}
	return Load(string(d))
}
