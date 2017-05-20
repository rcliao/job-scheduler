package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/ajvb/kala/client"
	"github.com/ajvb/kala/job"
)

func jobHandler(c *client.KalaClient) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createJob(c)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}

func main() {
	c := client.New("http://127.0.0.1:8000")
	http.HandleFunc("/job", jobHandler(c))
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func createJob(c *client.KalaClient) {
	// construct the current time string
	delay := rand.Int31n(60)
	t := time.Now().
		Add(time.Duration(delay) * time.Second).
		Format("2006-01-02T15:04:05-07:00")
	schedule := fmt.Sprintf("R2/%s/PT1S", t)
	fmt.Println(schedule)
	body := &job.Job{
		Schedule: schedule,
		Name:     "test_job",
		Command:  "curl 127.0.0.1:8080/hello",
	}
	id, err := c.CreateJob(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
