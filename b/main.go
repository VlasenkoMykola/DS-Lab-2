package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job func(in, out chan interface{})

func ExecuteJobs(jobs ...job) {
	in := make([]chan interface{}, 0, len(jobs))
	for i := 0; i < len(jobs); i++ {
		in = append(in, make(chan interface{}, 100))
	}

	wg := &sync.WaitGroup{}
	for i := range jobs {
		wg.Add(1)
		go func(counter int) {
			if counter != 0 {
				jobs[counter](in[counter-1], in[counter])
			} else {
				jobs[counter](in[counter], in[counter])
			}
			wg.Done()
			close(in[counter])
		}(i)
	}
	wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var property_value = [10]int{}
	var min int = 100
	var max int = 100000
	for i := 0; i < 10; i++ {
		property_value[i] = rand.Intn(max-min) + min
		fmt.Printf("created property worth $%v\n", property_value[i])
	}

	var storage job = func(in, out chan interface{}) {
		for i := 0; i < 10; i++ {
			out <- i
		}
	}
	var ivanov job = func(in, out chan interface{}) {
		for val := range in {
			time.Sleep(300 * time.Millisecond)
			out <- val
			fmt.Printf("Ivanov took property out from storage %v\n", val)
		}
	}
	var petrov job = func(in, out chan interface{}) {
		for val := range in {
			time.Sleep(200 * time.Millisecond)
			out <- val
			fmt.Printf("Petrov loaded property into a truck. %v\n", val)
		}
	}
	var nechiporchuk job = func(in, out chan interface{}) {
		sum := 0
		for val := range in {
			time.Sleep(100 * time.Millisecond)
			sum += property_value[val.(int)]
			fmt.Printf("Nechiporchuk added to the cost of the stolen %v property.\n", val)
		}
		fmt.Printf("Nechiporchuk counted $%v worth of property by the end\n", sum)
	}

	jobs := []job{
		storage,
		ivanov,
		petrov,
		nechiporchuk,
	}

	ExecuteJobs(jobs...)
}
