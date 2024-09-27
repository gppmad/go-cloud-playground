// package main

// // func main() {
// // 	jobs := make(chan int, 5)
// // 	done := make(chan bool)

// // 	go func() {
// // 		for {
// // 			j, more := <-jobs
// // 			if more {
// // 				fmt.Println("Received", j)
// // 			} else {
// // 				fmt.Println("All jobs have been received")
// // 				//done <- true
// // 				return
// // 			}

// // 		}

// // 	}()

// // 	for i := 0; i < 5; i++ {
// // 		jobs <- i
// // 		fmt.Println("Sent", i)
// // 	}
// // 	close(jobs)
// // 	<-done

// // }
