package git

import (
	"errors"
	"os/exec"
	"regexp"
)

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
