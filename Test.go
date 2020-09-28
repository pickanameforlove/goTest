package main

import (
	"sync"
)

var wg sync.WaitGroup

/*func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("Create Coroutines")
	go printPrime("A")
	go printPrime("B")
	fmt.Println("Waiting To Finish")
	//wg.Wait()
	fmt.Println("Terminate Program")
}

func printPrime(prefix string)  {
	defer wg.Done()
	next:
	for outer:=2;outer<5000;outer++{
		for inner := 2;inner<outer;inner++{
			if outer%inner==0{
				continue next
			}
		}
		fmt.Printf("%s:%d\n",prefix,outer)
	}
	fmt.Println("Completerd",prefix)

}*/
