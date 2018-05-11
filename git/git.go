package git

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func TemplateCommit() error {
	error := exec.Command("git", "checkout", "-b", "githum-template").Run()
	if error != nil {
		return error
	}
	if err := os.Mkdir(".github", 0755); err != nil {
		fmt.Println(err)
	}
	filename := filepath.Join(".github", "ISSUE_TEMPLATE.md")
	ioutil.WriteFile(filename, []byte("hello world!"), 0644)
	error = exec.Command("git", "add", ".github").Run()
	if error != nil {
		return error
	}
	error = exec.Command("git", "commit", "-m", "'add GitHub template'").Run()
	if error != nil {
		return error
	}
	error = exec.Command("git", "push", "origin", "github-template").Run()
	if error != nil {
		return error
	}
	return nil
}

func MainRemote() (string, error) {
	out, error := exec.Command("git", "remote", "-v").Output()
	if error != nil {
		return "", error
	}

	remotes := map[string]string{}
	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(out), -1) {
		matched := regexp.MustCompile(`(.+)\t.+github\.com:?(.+)\.git`).FindStringSubmatch(v)
		if len(matched) > 0 {
			remotes[matched[1]] = matched[2]
		}
	}
	if val, ok := remotes["upstream"]; ok {
		return val, nil
	}
	if val, ok := remotes["origin"]; ok {
		return val, nil
	}
	return "", errors.New("cloud not get the remote")
}
