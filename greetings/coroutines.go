package greetings

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func slowFunc(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "Sleeper Finished"

}


func slowFunc2(c chan string) {

	t := time.NewTicker(1 * time.Second)
   
    for {
   
   		c <- "ping"
   
   		<-t.C
   
    }
   
 }


func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}


func channelReader(messages <-chan string) {
    msg := <-messages
    fmt.Println(msg)
}

func channelWriter(messages chan<- string) {
    messages <- "Hello world"
}

func channelReaderAndWriter(messages chan string) {
     msg := <-messages
     fmt.Println(msg)
     messages <- "Hello world"
}



func runner(){

	c := make(chan string)

	go slowFunc(c)
	//fmt.Println("I am not shown until slofFunc completes")
	msg := <-c

	fmt.Printf("Msg from coroutine is %s", msg)
}

func responseTime(url string) {

	start := time.Now()
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	elapsed := time.Since(start).Seconds()

	fmt.Printf("%s took %v seconds \n", url, elapsed)

}


func channelRunner() {
	messages := make(chan string, 2)

	messages <- "Hello"
	messages <- "World"
	close(messages)

	fmt.Println("Pushed two messages onto Channel with no receivers")
	time.Sleep(time.Second * 1)
	receiver(messages)



}

func blockingRunner() {
	messages := make(chan string, 2)

	go slowFunc2(messages)
	
	for i := 0; i < 5; i++ {

		msg := <-messages
	
		fmt.Println(msg)
	
	}

}

func writerRunner() {

	c := make(chan string)

	go channelWriter(c)

	//fmt.Println("I am not shown until slofFunc completes")
	msg := <-c

	time.Sleep(3)

	fmt.Printf("Msg from coroutine is %s", msg)
}

func readerRunner() {

	c := make(chan string)

	go channelReader(c)
	
	c <- "Hello world"
	//fmt.Println("I am not shown until slofFunc completes")
	time.Sleep(3)

	
}

func readerWriterRunner() {

	c := make(chan string)

	go channelReaderAndWriter(c)
	
	c <- "Hello world from main"
	
	time.Sleep(3)
	
	msg := <-c

	//fmt.Println("I am not shown until slofFunc completes")
	
	fmt.Printf("Msg from coroutine is %s", msg)	
	
}






