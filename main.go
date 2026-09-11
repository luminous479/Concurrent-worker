package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	JobId int
}

func worker(workerId int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing Job %d\n", workerId, job.JobId)

		time.Sleep(time.Second)

		fmt.Printf("Worker %d finished Job %d\n", workerId, job.JobId)
	}
}

func main() {

	var wg sync.WaitGroup

	jobs := make(chan Job)

	jobList := []Job{
		{JobId: 1},
		{JobId: 2},
		{JobId: 3},
		{JobId: 4},
		{JobId: 5},
	}

	wg.Add(2)

	go worker(1, jobs, &wg)
	go worker(2, jobs, &wg)

	for _, job := range jobList {
		jobs <- job
	}

	close(jobs)

	wg.Wait()

	fmt.Println("All jobs completed")
}