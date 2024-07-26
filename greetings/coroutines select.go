package greetings

import (
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Second * 5)
	c <- "ping on channel 1"

}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "Ping on channel 2"
}


func sender(c chan string ) {

	t := time.NewTicker(1 * time.Second)

	for {
	
		c <- "I';m sending a message"
		<-t.C
	
	}

}




