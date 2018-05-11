package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/zaru/rep/git"
)

const apiHost = "api.github.com"
const userAgent = "Rep command agent"
const contentType = "application/json; charset=utf-8"
const githubApiVer = "application/vnd.github.symmetra-preview+json"

type Config struct {
	Labels []Label
}

type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
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
	res, err := api.PostJSON("POST", "/repos/"+remote+"/labels", label)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)
	fmt.Printf("%v", bodyString)
}

func (client *Client) api() *apiClient {
	httpClient := new()
	return &apiClient{
		httpClient:  httpClient,
		AccessToken: client.token(),
	}
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
