package main

import "time"

func main() {

	// 1. 심플 예제
	// ch := make(chan int) // 정수형 채널

	// go func() {
	// 	ch <- 123 // 채널에 123을 전달
	// }()

	// var i int
	// i = <- ch // 채널로부터 123을 받음
	// println(i)

	// 2. 대기 기능
	// done := make(chan bool)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	done <- true
	// }()

	// // True를 받을 때까지(=go로 선언된 루틴이 끝날 때까지) 대기. done에 true가 들어오면 그 때 값을 받고 끝난다.
	// <-done

	// 3. Buffered Channel vs Unbuffered Channel
	// Unbuffered Channel
	// c := make(chan int)
	// c <- 1           // 수신루틴이 없음. 이런 상황을 데드락이라고 한다. fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(<-c) // 별도의 Go루틴이 없으므로 데드락

	// // Buffered Channel
	// ch := make(chan int, 1)
	// ch <- 101
	// fmt.Println(<-ch)

	// 4. Channel Parameter
	// ch := make(chan string, 1)
	// sendChan(ch)
	// receiveChan(ch)

	// 5. 채널 닫기
	// ch := make(chan int, 2)

	// ch <- 1
	// ch <- 2

	// close(ch)

	// // 수신 argument <-ch는 2개의 리턴값을 갖는다 : 채널 메시지 & 성공 여부
	// println(<-ch)
	// println(<-ch)

	// if _, success := <-ch; !success {
	// 	println("No More Data")
	// }

	// 6. 채널 range문

	// ch := make(chan int, 2)

	// ch <- 1
	// ch <- 2

	// close(ch)

	// // 수신자가 임의의 갯수를 데이터를 수신하는 2가지 방법
	// // 1. 채널 닫힌 걸 감지할 때 까지 수신 : 무한 for문에서 2번째 파라미터 확인
	// for {
	// 	if i, success := <-ch; success {
	// 		println(i)
	// 	} else {
	// 		break
	// 	}
	// }

	// // 2. 위 표현과 동일함 : 동일한 방식, for range문으로 간결하게 표현
	// for i := range ch {
	// 	println(i)
	// }

	// 7. 채널 Select문
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("Run1 완료")
		case <-done2:
			println("Run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

// func sendChan(ch chan<- string) { // 송신 파라미터 chan<-
// 	ch <- "Data"
// 	// x := <- ch : 수신 시도 - 에러 발생
// }

// func receiveChan(ch <-chan string) { // 수신 파라미터 <-chan
// 	data := <-ch
// 	fmt.Println(data)
// }

//
