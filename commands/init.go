package commands

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/zaru/rep/client"
)

func Init(filePath string) {
	var config client.Config
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	client := client.NewClient()
	for _, v := range config.Labels {
		client.AddLabel(v)
	}
	client.AddFile("ISSUE_TEMPLATE.md", config.Issue.Template)
	client.AddFile("PULL_REQUEST_TEMPLATE.md", config.PullRequest.Template)
}
