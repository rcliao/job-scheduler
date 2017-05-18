package main

import (
	"fmt"
	"net/http"

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
	body := &job.Job{
		Schedule: "R2/2017-05-18T01:31:00.828696-07:00/PT10S",
		Name:     "test_job",
		Command:  "curl 127.0.0.1:8080/hello",
	}
	id, err := c.CreateJob(body)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
