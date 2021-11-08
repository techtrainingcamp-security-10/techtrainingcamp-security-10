package resource

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type FullConfig struct {
	Redis RedisOpts `yml:"redis"`
	Mysql MySQLOpts `yml:"mysql"`
}

func GetConfig() (FullConfig, error) {
	config := FullConfig{}
	pwd, _ := os.Getwd()
	cfg, err := ioutil.ReadFile(path.Join(pwd, `/internal/resource/config.yml`))
	if err != nil {
		return config, err
	}
	err = yaml.UnmarshalStrict(cfg, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
