# Go Worker Pool

A simple Worker Pool implementation in Go using **Goroutines, Channels, and WaitGroups**.

## What I Learned

- Goroutines for concurrent execution
- Channels for passing jobs between workers
- Worker Pool pattern
- `sync.WaitGroup` for synchronization
- Closing channels after sending jobs
- Processing multiple jobs concurrently

## How It Works

The program creates a fixed number of workers.  
Jobs are sent through a channel, and available workers process them concurrently.

```text
Main
 |
 |--- Job 1
 |--- Job 2
 |--- Job 3
 |--- Job 4
 |--- Job 5
       |
     Channel
       |
   -------------
   |     |     |
 Worker Worker Worker
   1     2     3