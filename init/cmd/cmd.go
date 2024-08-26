package cmd

import (
	"Back/controllers"
	"Back/repositories"
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	MySQL struct {
		Host string
	}
	Mongo struct {
		URI string
	}
}

func NewCmd(filePath string) {
	c := new(Config)
	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	sql, nsql, err := repositories.NewRepository(c.MySQL.Host, c.Mongo.URI)
	if err != nil {
		panic(err)
	}
	err = controllers.NewContorllers(c.Server.Port, sql, nsql)
	if err != nil {
		panic(err)
	}
}
