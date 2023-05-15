package projects

import (
	_ "embed"
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

//go:embed config.hcl
var configData []byte

type Project struct {
	Slug        string `hcl:"slug,label"`
	Name        string `hcl:"name"`
	Description string `hcl:"description"`
}

type Config struct {
	ProjectList []Project `hcl:"project,block"`
}

func Get() (Config, error) {
	config := Config{}
	err := hclsimple.Decode("config.hcl", configData, nil, &config)
	if err != nil {
		return config, fmt.Errorf("decoding project config from config.hcl: %w", err)
	}
	return config, nil
}
