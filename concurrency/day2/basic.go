package main

import ()

// func worker(id int) {
// 	fmt.Println("Worker starting...!")
// 	fmt.Printf("Worker  :%d\n", id)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Worker done")
// }

// worker-pool concept
// func work(id int, jobs <-chan int, results chan<- int) {
// 	fmt.Println("Starting working go routine function")
// 	for j := range jobs {
// 		fmt.Printf("Starting worker id : %d job  :%d \n", id, j)
// 		time.Sleep(1 * time.Second)
// 		fmt.Printf("job done worker id : %d job  :%d \n", id, j)
// 		results <- 2 * j
// 	}
// }

func main() {
	// fmt.Println("Initial goroutine ", runtime.NumGoroutine())

	// for i := 0; i < 199; i++ {
	// 	go worker(i)
	// }

	// fmt.Println("finished goroutines", runtime.NumGoroutine())
	// time.Sleep(5 * time.Second)

	// workChh := make(chan int)
	// resultChh := make(chan int)

	// numWorker := 3

	// // starting worker here
	// for i := 1; i <= 3; i++ {
	// 	go work(i, workChh, resultChh)
	// }

	// // assign job of worker
	// for j := 1; j <= numWorker; j++ {
	// 	workChh <- j
	// }
	// close(workChh)

	// calling other function
	//Mutex()

	// example
	exa1()
}
