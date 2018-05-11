package commands

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/zaru/rep/client"
)

func Init() {
	var config client.Config
	_, err := toml.DecodeFile("./config.sample.toml", &config)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Printf("%v", config)

	client := client.NewClient()
	for _, v := range config.Labels {
		client.AddLabel(v)
	}
}
