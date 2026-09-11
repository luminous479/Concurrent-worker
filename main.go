package main

import (
	"fmt"
	"sync"
	"time"
)


func worker(workerId int, job Job, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("worker %d progressing Job %d\n",workerId, job.JobId )

     time.Sleep(time.Second)

    fmt.Printf("worker %d finished Job %d\n",workerId, job.JobId )
}

type Job struct{
	JobId int 
}

func main() {

 var wg sync.WaitGroup

  jobs := []Job {

	Job{JobId: 1},
	Job{JobId: 2},
	Job{JobId: 3},
	Job{JobId: 4},
	Job{JobId: 5},

  }

  wg.Add(len(jobs))

  for i, job := range jobs{
	go worker( i+1 , job, &wg) 
  }

  wg.Wait()

}