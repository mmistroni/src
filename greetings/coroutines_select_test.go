package greetings

import (
	"fmt"
	"testing"
	"time"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestRunningSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	select {
		case msg := <- channel1:
			fmt.Println("Received:" + msg)
		case msg := <- channel2:
			fmt.Println("Received:" + msg)

	}

}

func TestRunningSelectWithTimeout(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	select {
		case msg := <- channel1:
			fmt.Println("Received:" + msg)
		case msg := <- channel2:
			fmt.Println("Received:" + msg)
		case <- time.After(500 * time.Millisecond):
			fmt.Println("No message received after 500ms")

	}

}

