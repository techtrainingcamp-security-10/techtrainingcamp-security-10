package resource

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
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
		config = FullConfig{
			Redis: RedisOpts{
				Host: "127.0.0.1:6379",
			},
			Mysql: MySQLOpts{
				User:     "root",
				Password: "admin",
				Address:  "127.0.0.1:3306",
				Name:     "techtrainingcamp",
			},
		}
		return config, nil
	}
	err = yaml.UnmarshalStrict(cfg, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
