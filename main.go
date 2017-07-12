package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/ajvb/kala/client"
	"github.com/ajvb/kala/job"
)

// Manager is the basic unit to represent consumer to maintain their own worker
type Manager struct {
	ID      string
	Name    string
	Workers []Worker
}

// Worker represents the base simulation unit with efficiency
type Worker struct {
	Manager      Manager
	Efficienct   int
	Name         string
	CurrentJobID string
}

func main() {
	// TODO: get the schedudler server url from environment variable
	c := client.New("http://127.0.0.1:8000")

	// TODO: maintain a list of resources

	// TODO: create a list of workers each with different efficiency

	http.HandleFunc("/job", jobHandler(c))
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func jobHandler(c *client.KalaClient) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: parse the callback URL
		url := "127.0.0.1:8080/hello"
		name := "test_job"
		createJob(c, url, name)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello job simulation server")
	r.Println("Hello job simulation server")
}

func createJob(c *client.KalaClient, url, name string) {
	// construct the current time string
	delay := rand.Int31n(60)
	t := time.Now().
		Add(time.Duration(delay) * time.Second).
		Format("2006-01-02T15:04:05-07:00")
	schedule := fmt.Sprintf("R2/%s/PT1S", t)
	body := &job.Job{
		Schedule: schedule,
		Name:     name,
		Command:  "curl " + url,
	}
	id, err := c.CreateJob(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
