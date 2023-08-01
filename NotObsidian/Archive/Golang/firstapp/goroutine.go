package main

import (
	"fmt"
	"sync"
)

// func say(s string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(s, "***", i)
// 	}
// }

// func main() {
// 	//함수 동기적 실행
// 	say("Sync")

// 	//함수 비동기적 실행
// 	go say("Async1")
// 	go say("Async2")
// 	go say("Async3")

// 	// 3초 대기
// 	time.Sleep(time.Second * 3)
// }

// 2. 익명함수 go루틴
// go 뒤에서 선언되는 익명함수는 비동기로 실행된다.
func main() {
	var wait sync.WaitGroup // Goroutine을 기다리는 역할
	wait.Add(2)             // 2개

	go func() {
		defer wait.Done() // Done은 끝났다는 신호 전달함
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait()
}

// 3. Go는 다중 CPU처리를 켤 수 있다.
package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(4) // 4개의 cpu를 켜겠다는 뜻
}
