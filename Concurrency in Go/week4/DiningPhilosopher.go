package main

import (
	"fmt"
	"sync"
	"time" 
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.
1.	There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2.	Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3.	The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4.	In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5.	The host allows no more than 2 philosophers to eat concurrently.
6.	Each philosopher is numbered, 1 through 5.
7.	When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
8.	When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
var wg sync.WaitGroup 
type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	position        int
	times           int
	finishToEat       chan int
	startToEat        chan int
	reqToEat          chan int
}

func (p Philo) eat() {
	for p.times < 3 {  
		p.reqToEat <- p.position

		if eat := <- p.startToEat; eat == 1{
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Println("starting to eat ", p.position)
			time.Sleep(1000)
			fmt.Println("finisihing to eat ", p.position)
			p.times ++
			p.finishToEat <- p.position
			p.rightCS.Unlock()
			p.leftCS.Unlock() 

		}
	}  
	wg.Done()
	return
 
}

//5.	The host allows no more than 2 philosophers to eat concurrently.
func askToTheHost(reqToEat, start, finish, quitToEat, quitFromHost chan int) {
	spots := make(map[int]string)
 
	for {
		select {
		case position:= <-reqToEat:
			if len(spots) < 2 {
				spots[position] = "Eating"
				start <- 1 
			} else {
				///waiting for the turn
				start <- 0 
			}
		case position := <- finish:
			delete(spots, position) 
		 }
	} 
	
}


func main() {
	finishToEat := make(chan int)
	startToEat := make(chan int)
	requestEat := make(chan int)
	quitFromHost := make(chan int)
	quitToEat := make(chan int)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1, 0, finishToEat, startToEat, requestEat}
	}

	wg.Add(5)
	//4.	In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
	go askToTheHost(requestEat, startToEat, finishToEat, quitToEat, quitFromHost)
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

	wg.Wait()
	fmt.Println("Dining Philosopher finished!!")

}
