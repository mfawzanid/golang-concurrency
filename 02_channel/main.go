package main

import (
	"fmt"
	"time"
)

/*
Channels are typed communication pipelines that allow goroutines to safely and synchronously exchange data.
They provide a way to send and receive values of a specific type, enabling coordination and data sharing between concurrent operations.

We simulate the use of a channel by modeling a typical order processing flow.
Once an order is completed, it triggers a series of follow up notifications.
These notifications are not critical to the main business flow and can be handled asynchronously.
*/

// Simulate order validation and sends the validated user into the channel
func processOrder(c chan string, user string) {
	fmt.Printf("[🔍] Validating order for %v...\n", user)
	fmt.Printf("[✅] Finished validating order for %v\n", user)
	c <- user
}

// Simulate notification mechanisms
func sendEmail(user string) {
	fmt.Printf("📧 Sending email to %v...\n", user)
}

func sendPushNotification(user string) {
	fmt.Printf("🔔 Sending push notification to %v...\n", user)
}

func sendSMS(user string) {
	fmt.Printf("📱 Sending SMS to %v...\n", user)
}

func main() {
	users := []string{"adi", "budi", "caca", "dedi"}

	// Channel used to communicate validated orders to the notification dispatcher
	c := make(chan string)

	// Notification dispatcher: receives validated orders and spawns workers
	go func() {
		for {
			fmt.Printf("[📦] Waiting for completed order at %v...\n", time.Now().Format("15:04:05.00000"))
			user := <-c
			fmt.Printf("[📬] Dispatching notifications for %v at %v...\n", user, time.Now().Format("15:04:05.00000"))

			// Send all notification types concurrently
			go sendEmail(user)
			go sendPushNotification(user)
			go sendSMS(user)
		}
	}()

	// Simulate order validation pipeline
	for i, user := range users {
		processOrder(c, user)

		// Simulate batch delay (after every 2 users)
		if (i+1)%2 == 0 {
			fmt.Println("[⏳] Sleeping between batches...")
			time.Sleep(time.Second)
		}
	}

	// Allow time for goroutines to finish (demo purpose only)
	time.Sleep(1 * time.Second)
}
