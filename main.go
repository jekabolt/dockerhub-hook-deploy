package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"log"

	"net/http"

	"github.com/caarlos0/env"
	"github.com/go-chi/chi"
)

type Client struct {
	Port  string `env:"SERVER_PORT" envDefault:"8555"`
	Token string `env:"DOCKER_AUTH_TOKEN" envDefault:"kektoken"`
}

type DockerHook struct {
	CallbackURL string `json:"callback_url"`
	PushData    struct {
		Images []string `json:"images"`
		Pusher string   `json:"pusher"`
		Tag    string   `json:"tag"`
	} `json:"push_data"`
	Repository struct {
		CommentCount int    `json:"comment_count"`
		Description  string `json:"description"`
		Name         string `json:"name"`
		Namespace    string `json:"namespace"`
		Owner        string `json:"owner"`
		RepoName     string `json:"repo_name"`
		RepoURL      string `json:"repo_url"`
	} `json:"repository"`
}

func main() {

	c := &Client{}
	err := env.Parse(c)
	if err != nil {
		log.Fatalf("main:env.Parse [%v]", err.Error())
	}

	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.Logger)
	fmt.Println("--- ", fmt.Sprintf("/%s", c.Token))
	r.Post(fmt.Sprintf("/%s", c.Token), c.hookHandler)

	block := make(chan string)
	go func() {
		http.ListenAndServe(":"+c.Port, r)
	}()
	<-block
}

func (c *Client) hookHandler(w http.ResponseWriter, req *http.Request) {
	var dh DockerHook
	err := json.NewDecoder(req.Body).Decode(&dh)
	if err != nil {
		log.Printf("hookHandler:json.NewDecoder [%v]", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(dh)

	if strings.Contains(dh.PushData.Tag, "release") {

	}
	// 	docker pull ${TRAVIS_REPO_SLUG}:${IMAGE_VERSION}
	// docker run -d ${TRAVIS_REPO_SLUG}:${IMAGE_VERSION}
	app := "docker"
	pullCmd := exec.Command(app, "pull", dh.Repository.RepoName+":"+dh.PushData.Tag)

	runCmd := exec.Command(app, "run", "-d", dh.Repository.RepoName+":"+dh.PushData.Tag)

	output, err := pullCmd.Output()
	fmt.Printf("output  %s", output)

	output2, err := runCmd.Output()
	fmt.Printf("output2 %s ", output2)

}
