/*
 =================================================================
 FILE...............: sprintRelay.go
 DESCRIPTION........: Program that simulates a relay race in which
                      a team has four runners. The second, third
				      and fourth runners can not start running
                      until they receive the baton delivered by
 					  the runner who preceded it.
 AUTHOR.............: Luís Eduardo (cruxiu@ufrn.edu.br)
 CREATED IN.........: 19/06/2018
 MODIFIED IN........: 21/06/2018
 =================================================================
*/
package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

// Struct representing a runner
type Runner struct {
	name string
}

// Struct representing a team of runners
type Team struct {
	country string
	time int
	runners []Runner
}

/* Function that returns the index of a Runner object in a
slice passed with parameter */
func index(slice []Runner, runner Runner) int {
	for i, _ := range slice {
		if slice[i] == runner {
			return i
		}
	}
	return -1
}

/* Function that generates a random integer between a minimum
and maximum number passed as a parameter. */
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

/* Function that will make the runner run for a while. The next
runner will be waiting for the current runner until he finishes
the route. */
func run(runner Runner, team *Team, wg *sync.WaitGroup, m *sync.Mutex) {
	fmt.Printf("The runner %s of %s started waiting for the baton. \n", runner.name, team.country)
	m.Lock()

	position := index(team.runners, runner)
	if position < len(team.runners) - 1 {
		go run(team.runners[position+1], team, wg, m)
	}

	fmt.Printf("The runner %s of %s started running with the baton. \n", runner.name, team.country)
	duration := random(3, 12)
	team.time += duration
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("The runner %s of %s finished his run and took %d seconds on the route. \n", runner.name, team.country, duration)

	m.Unlock()
	wg.Done()
}

// Main function
func main() {

	// The team of runners
	jamaicaTeam := Team{"Jamaica", 0, []Runner{
													{"Bolt"},
													{"Blake"},
													{"Powell"},
													{"Ashmeade"}}}

	fmt.Println("Welcome to the 4x100 metres relay simulator!")

	fmt.Println("The jamaica team's athletes will be:")
	for index, runner := range jamaicaTeam.runners {
		fmt.Println("Runner", index + 1, runner.name)
	}

	fmt.Println("The race will start! \n")

	// Waiting group for collective synchronization of the runners
	var wg sync.WaitGroup
	wg.Add(4)

	// Mutex for synchronization by mutual exclusion between runners
	var m sync.Mutex

	// Call the function that will make the first runner prepare to start running
	run(jamaicaTeam.runners[0], &jamaicaTeam, &wg, &m)

	//Wait all runners
	wg.Wait()

	fmt.Println("The race is over :( \n")

	fmt.Printf("The team of %s took %d seconds to complete the race.", jamaicaTeam.country, jamaicaTeam.time)
}