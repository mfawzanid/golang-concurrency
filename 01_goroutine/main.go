package main

import (
	"fmt"
	"time"
)

// Goroutines official definition
// A goroutine is a lightweight thread managed by the Go runtime.

/*
	We simulate the use of a goroutine by modeling a typical payment process.
	In real-world systems, updating the payment status often includes logging the transaction,
	but logging is not critical to the main flow and can be performed asynchronously.

	Therefore, in this simulation, we run the logging process in a separate goroutine
	to represent background tasks that do not need to block the main operation.
*/

const (
	timeSleep = 100 // milliseconds
)

// simulate logging transaction in multiple steps
func logTransaction() {
	totalSteps := 5

	for i := 1; i <= totalSteps; i++ {
		fmt.Printf("Logging transaction... [step %d]\n", i)

		if i == totalSteps {
			fmt.Printf("[done] Logging transaction\n")
		}

		time.Sleep(time.Duration(timeSleep) * time.Millisecond)
	}
}

// simulate updating payment status in multiple steps
func updatePaymentStatus() {
	totalSteps := 5 // Note: decrease this value to simulate Question 1 below

	for i := 1; i <= totalSteps; i++ {
		fmt.Printf("Updating payment status... [step %d]\n", i)

		if i == totalSteps {
			fmt.Printf("[done] Updating payment status\n")
		}

		time.Sleep(time.Duration(timeSleep) * time.Millisecond)
	}
}

func processPayment() {
	go logTransaction() // Run in background

	// Note: Comment out the line below to simulate Question 2 below
	updatePaymentStatus() // Run in main goroutine
}

func main() {
	processPayment()

	/*
		Explanation:
		- A goroutine runs asynchronously and is managed by the Go runtime.
		- The `main` function is also a goroutine. When it finishes, it terminates the entire program,
		  even if other goroutines are still running.

		Question 1:
		What happens if `updatePaymentStatus` finishes before `logTransaction`?
		> You can simulate this by setting a smaller `totalSteps` for `updatePaymentStatus`.
		> Since `updatePaymentStatus()` runs in the main goroutine, once it finishes,
		  the `main()` function ends and so the background goroutine (`logTransaction`) is killed early.

		Question 2:
		What happens if we remove `updatePaymentStatus()` and only run `go logTransaction()`?
		> The `main()` function will exit immediately, and the goroutine may never run at all,
		  as Go does not wait for any goroutine by default.
	*/
}
