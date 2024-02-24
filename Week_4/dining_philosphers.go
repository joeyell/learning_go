package week_4

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	eating          int
	number          int
	leftCS, rightCS *ChopS
}

type Host struct {
	eatingTracker     int
	permissionRequest chan bool
	answerPermission  chan bool
	eatingEvents      chan int
	sync.Mutex
}

func (p *Philo) eat(host *Host, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(host.eatingTracker)
		p.eating++
		host.Lock()
		p.getPermission(host)
		host.Unlock()
		//fmt.Printf("Philosopher %d has started meal %d\n", p.number, p.eating)
		p.removePermission(host)
		//fmt.Printf("Philosopher %d has finished eating %d\n", p.number, p.eating)

	}
}

func (host *Host) bouncer() {
	for i := range host.permissionRequest {
		if host.eatingTracker > 2 {
			fmt.Println(host.eatingTracker)
		}
		switch i {
		case true:
			if host.eatingTracker < 2 {
				host.answerPermission <- true // permission granted
				host.eatingTracker++
			} else {
				host.answerPermission <- false // permission not granted
			}
		case false:
			host.eatingTracker--
		}

	}
}

func (p *Philo) getPermission(host *Host) {
	//host.Lock()
	host.permissionRequest <- true //ask permission
	i := <-host.answerPermission   //wait for host approval
	if i {
		//fmt.Println("Permission received")
		p.leftCS.Lock()  //lock chopsticks
		p.rightCS.Lock() //
		//host.eatingEvents <- 1
	} else {
		p.getPermission(host)
	}
	//host.Unlock()
}

func (p *Philo) removePermission(host *Host) {
	host.permissionRequest <- false //revoke permission
	//fmt.Println("Permission revoked")
	p.leftCS.Unlock()  //Unlock chopsticks
	p.rightCS.Unlock() //
	//host.eatingEvents <- -1
}

func DiningPhilosophers() {
	philosophers := 20
	// initialise chopsticks
	CSticks := make([]*ChopS, philosophers)
	for i := 0; i < philosophers; i++ {
		CSticks[i] = new(ChopS)
	}

	// initialise philosphers
	philos := make([]*Philo, philosophers)
	for i := 0; i < philosophers; i++ {
		philos[i] = &Philo{0, i + 1, CSticks[i], CSticks[(i+1)%5]} // % to handle philospher 5's final chopstick
	}

	// initialise host
	var philoHost Host
	philoHost.permissionRequest = make(chan bool)
	philoHost.answerPermission = make(chan bool)
	philoHost.eatingEvents = make(chan int, philosophers*3)

	// Start eating
	var wg sync.WaitGroup
	wg.Add(philosophers)

	go philoHost.bouncer()

	for i := 0; i < philosophers; i++ {
		go philos[i].eat(&philoHost, &wg)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(philoHost.permissionRequest) // Close the channel after all goroutines are done
		close(philoHost.answerPermission)
		close(philoHost.eatingEvents)
	}()

	// // Count the number of philosophers eating simultaneously
	// count := 0
	// for event := range philoHost.eatingEvents {
	// 	if event == 1 { // eating start event
	// 		count++
	// 	} else if event == -1 { // eating end event
	// 		count--
	// 	}
	// 	if count >= 3 {
	// 		fmt.Println("ILLEGAL", count)
	// 	}
	// }
	wg.Wait()
}
