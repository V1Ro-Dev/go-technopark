package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

type WorkerPool struct {
	workers map[uuid.UUID]chan struct{}
	mu      sync.Mutex
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		workers: make(map[uuid.UUID]chan struct{}),
	}
}

func (wp *WorkerPool) startWorker(in <-chan string, quit <-chan struct{}) {
	for {
		select {
		case input, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(input)
		case <-quit:
			return
		}
	}
}

func (wp *WorkerPool) AddWorker(in <-chan string) {
	quit := make(chan struct{})
	id := uuid.New()

	wp.mu.Lock()
	wp.workers[id] = quit
	wp.mu.Unlock()

	go wp.startWorker(in, quit)
	fmt.Printf("Worker %s added\n", id)
}

func (wp *WorkerPool) RemoveWorker() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	for id, quit := range wp.workers {
		close(quit)
		delete(wp.workers, id)
		fmt.Printf("Worker %s removed\n", id)
		return
	}
	fmt.Println("No workers to remove")
}

func main() {
	wp := NewWorkerPool()
	in := make(chan string)

	for i := 0; i < 3; i++ {
		wp.AddWorker(in)
	}

	var testData = make([]string, 1000)
	for i := 0; i < 1000; i++ {
		testData[i] = strconv.Itoa(i)
	}
	for _, month := range testData {
		in <- month
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Commands:")
	fmt.Println("  add    - add a worker")
	fmt.Println("  remove - remove a worker")
	fmt.Println("  exit   - quit program")
	fmt.Println("Type any other text to send it as a task to the pool.")

	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "add":
			wp.AddWorker(in)
		case "remove":
			wp.RemoveWorker()
		case "exit":
			fmt.Println("Exiting...")
			close(in)
			return
		default:
			in <- text
		}
	}
}
