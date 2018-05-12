package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zaru/rep/git"
)

const apiHost = "api.github.com"
const userAgent = "Rep command agent"
const contentType = "application/json; charset=utf-8"
const githubApiVer = "application/vnd.github.symmetra-preview+json"

type Config struct {
	Labels      []Label
	Issue       Template
	PullRequest Template `toml:"pull_request"`
}

type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type Template struct {
	Template string
}

type TemplateFile struct {
	Message string `json:"message"`
	Content string `json:"content"`
	Sha     string `json:"sha"`
}

type apiClient struct {
	httpClient  *http.Client
	AccessToken string
}
type apiResponse struct {
	*http.Response
}
type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func new() *http.Client {
	return &http.Client{}
}

func (client *Client) AddLabel(label Label) {
	api := client.api()
	remote, _ := git.MainRemote()

	exists := client.LabelExists(label.Name)
	method := "POST"
	url := "/repos/" + remote + "/labels"
	if exists {
		method = "PATCH"
		url += "/" + label.Name
	}

	_, err := api.PostJSON(method, url, label)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
}

func (client *Client) AddFile(name, template string) {
	api := client.api()
	remote, _ := git.MainRemote()

	sha := client.GetShaOfFile(name)

	file := TemplateFile{
		Message: "add GitHub template file",
		Content: base64.StdEncoding.EncodeToString([]byte(template)),
		Sha:     sha,
	}

	_, err := api.PostJSON("PUT", "/repos/"+remote+"/contents/.github/"+name, file)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
}

func (client *Client) LabelExists(name string) bool {
	res := client.GetLabel(name)
	if res == nil {
		return false
	}
	return true
}

func (client *Client) GetLabel(name string) *apiResponse {
	api := client.api()
	remote, _ := git.MainRemote()
	res, err := api.Get("/repos/" + remote + "/labels/" + name)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	if res.StatusCode == 404 {
		return nil
	}
	return res
}

func (client *Client) GetShaOfFile(name string) string {
	res := client.GetFile(name)
	if res == nil {
		return ""
	}
	var result struct {
		Sha string `json:"sha"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		fmt.Errorf("Error: %v", err)
	}
	return result.Sha
}

func (client *Client) GetFile(name string) *apiResponse {
	api := client.api()
	remote, _ := git.MainRemote()
	res, err := api.Get("/repos/" + remote + "/contents/.github/" + name)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	if res.StatusCode == 404 {
		return nil
	}
	return res
}

func (client *Client) api() *apiClient {
	httpClient := new()
	return &apiClient{
		httpClient:  httpClient,
		AccessToken: client.token(),
	}
}

func (api *apiClient) Get(url string) (res *apiResponse, err error) {
	req, err := http.NewRequest("GET", "https://"+apiHost+url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", "token "+api.AccessToken)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", githubApiVer)
	httpResponse, err := api.httpClient.Do(req)
	if err != nil {
		return
	}

	res = &apiResponse{httpResponse}
	return
}

func (api *apiClient) PostJSON(method, url string, body interface{}) (res *apiResponse, err error) {
	json, err := json.Marshal(body)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(json)
	req, err := http.NewRequest(method, "https://"+apiHost+url, buf)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", "token "+api.AccessToken)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", githubApiVer)
	httpResponse, err := api.httpClient.Do(req)
	if err != nil {
		return
	}

	res = &apiResponse{httpResponse}
	return
}

func (client *Client) token() string {
	token := os.Getenv("GITHUB_TOKEN")
	return token
}
