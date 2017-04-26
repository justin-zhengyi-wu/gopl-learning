package main

import (
	"fmt"
	"time"
)

func main() {
	// hello()
	// values()
	// a := "asdf"
	// ret := whatAmI(a)
	// fmt.Println(ret)
	//	_arrays()
	//	_slices()
	//	_range()
	//	pointer()
	/*
		s := Person{name: "Justin", age: 30}
		s2 := &s
		s2.age = 31
		fmt.Println(s.age)
	*/
	//	_goroutines()
	//	_channels()
	//	channelBuffer()
	//	channelSync()
	//	channelDirection()
	//	_select()
	workerPools()
}

type Person struct {
	name string
	age  int
}

func hello() {
	fmt.Println("Hello world!")
}

func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func _switch() {
	i := 2
	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	}
}

func whatAmI(i interface{}) string {
	switch t := i.(type) {
	case bool:
		return "bool"
	case int:
		return "int"
	default:
		return fmt.Sprintf("%T", t)
	}
}

func _arrays() {
	var a [5]int
	fmt.Println(a)
	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])
}

func _slices() {
	a := []int{}
	b := make([]int, 5)
	b[1] = 1
	c := make([]string, 3)
	fmt.Println(a)
	fmt.Println(b, c)
}

func _range() {
	for i, c := range "ABC" {
		fmt.Println(i, c)
	}
}

func pointer() {
	i := 1
	fmt.Println(i)
	func(j int) {
		j = 2
	}(i)
	fmt.Println(i)
	func(k *int) {
		*k = 3
	}(&i)
	fmt.Println(i)
	fmt.Println(&i)
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func _goroutines() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("xxx")
	fmt.Println("Print")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}

func _channels() {
	msg := make(chan string)
	fmt.Println("1")
	go func() { msg <- "ping" }()
	fmt.Println(msg)
	msg2 := <-msg
	fmt.Println(msg2)
	fmt.Println(msg, msg2)
}

func channelBuffer() {
	msg := make(chan string, 2)
	msg <- "msg1"
	msg <- "msg2"
	fmt.Println(<-msg)
	fmt.Println(<-msg)
	//	fmt.Println(<-msg)
}

func channelSync() {
	w := make(chan bool, 1)
	go worker(w)
	ret := <-w
	fmt.Println(ret)

}
func worker(w chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("one")
	w <- true
}

func channelDirection() {
	chout := make(chan string, 1)
	chin := make(chan string, 1)
	ping(chin, "hi")
	//	fmt.Println(<-chin)
	pong(chin, chout)
	fmt.Println(<-chout)
}
func ping(ch chan<- string, msg string) {
	ch <- msg
}
func pong(chout <-chan string, chin chan<- string) {
	msg := <-chout
	chin <- msg
}

func _select() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("got", msg1)
		case msg2 := <-c2:
			fmt.Println("got", msg2)
		}
	}
}

func workerPools() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		<-results
	}
}
func worker2(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Println("worker", id, "started job", i)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", i)
		results <- i * 2
	}
}
