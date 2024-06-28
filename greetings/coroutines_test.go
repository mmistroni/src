package greetings

import (
	"testing"
	"time"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestRunning(t *testing.T) {
	runner()

}

func TestResponseTime(t *testing.T) {
	urls := make([]string, 3)

	urls[0] = "https://www.usa.gov/"
	urls[1] = "https://www.gov.uk/"
	urls[2] = "http://www.gouvernement.fr/"
	
	for _, u := range urls {
		go responseTime(u)
	}
	
	time.Sleep(time.Second * 5)
}

func TestChannelRunner(t *testing.T) {
	channelRunner()
}


func TestBlockingRunner(t *testing.T) {
	blockingRunner()
}

func TestChannelWriter(t *testing.T) {
	writerRunner()

}

func TestChannelReader(t *testing.T) {
	readerRunner()

}

func TestChannelReaderWriter(t *testing.T) {
	readerWriterRunner()

}

