package commands

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Labels []Label
}

type Label struct {
	Name  string
	Color string
}

func Init() {
	var config Config
	_, err := toml.DecodeFile("./config.sample.toml", &config)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Printf("%v", config)
}
