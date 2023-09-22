// DiningPhilosopers project main.go
package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type PhiloMessage struct {
	requestChannel chan int
}

type Philo struct {
	leftCS, rightCS *ChopS
	wg              *sync.WaitGroup
	id              int
	perm            chan int
	done            chan int
}

func host(permissions, done chan int) {
	fmt.Println("Host placing 2 permissions...")
	permissions <- 1
	permissions <- 1
	for {
		select {
		case result := <-done:
			fmt.Printf("Received  % v done signal. Adding permissions\n", result)
			permissions <- result
		default:
			fmt.Println("waiting..")
		}
	}

}

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		// requesting permissions
		ready := <-p.perm
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat %v \n", p.id)
		fmt.Printf("eating %v \n", i)
		fmt.Printf("finished eating %v \n", p.id)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		// notify host that you are done
		p.done <- ready
	}
	p.wg.Done()

}

func main() {
	fmt.Println("Initializing....!")

	var wg sync.WaitGroup

	var doneChannel = make(chan int)
	var permissionsChannel = make(chan int, 2)

	go host(permissionsChannel, doneChannel)

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i],
			CSticks[(i+1)%5], &wg, i, permissionsChannel, doneChannel}

	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat()
	}

	//Wait
	wg.Wait()
	print("Done")
}
