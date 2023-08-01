package main

import (
	"fmt"
	"strconv"
)

// 변수를 설정하는 다른 방법. main() 밖에서 할 수 있음
// 이를 패키지 레벨이라고 함
// var i int = 42
// 1번째 방법은 var i int까지는 가능
// 이 경우 초기화는 main 함수 내에서 해야 한다
// 3번째 방법인 :=는 여기서 사용 불가능하다.

// 패키지 레벨에서 할 수 있는 또 다른 일
// var actor string = "Elisabeth Sladen"
// var companion string = "Sarah Jane Smith"
// var doctorNumber int = 3
// var season int = 11
// 하나하나 var로 선언하지 않고, 묶을 수 있음
// var (
// 	actor        string = "Elisabeth Sladen"
// 	companion    string = "Sarah Jane Smith"
// 	doctorNumber int    = 3
// 	season       int    = 11
// )

// Shadowing - 가장 안쪽의 변수 선언이 가장 우선시된다
// var i int = 27

func main() {
	// fmt.Println(i)
	// var i int = 42
	// fmt.Println(i)
	// 방법 1 : 변수 선언 but 초기화 x
	// var i int // int i 가 아니다
	// i = 42
	// var i float32 = "foo"
	// 방법 2 : 변수 선언 & 초기화
	// var j float32 = 27
	// fmt.Printf("%v, %T", j, j) // printf는 포맷팅을 제공한다. ln과 달리
	// %v : 구조체 인스턴스 출력
	// %T : 구조체의 타입 출력
	// 방법 3
	// k := 42.0
	// fmt.Printf("%v, %T", k, k)
	// fmt.Println(i, j)
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
