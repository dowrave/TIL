package main

import (
	"fmt"
)

func main() {
	// 1. 반복문 기본 틀은 c와 동일함

	// 작동방식 : 0 출력 -> i++로 1 -> 1 출력 -> else로 3 -> i++로 4 -> 4 출력 -> if 문으로 2 -> 2+1 = 3 -> 3 출력 -> else로 7 -> 반복문 종료
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i%2 == 0 {
	// 		i /= 2 // i = i / 2
	// 	} else {
	// 		i = 2*i + 1
	// 	}
	// }
	// Counter (i)를 반복문 안에 넣는 건 좋은 예시가 아님. 하지 말라고 보여주는 거임

	// 여러 변수를 동시에 돌리고 싶어요
	// for i := 0 , j := 0; i < 5; i++, j++  //이 방식은 불가능 : go의 Increment는 Increment가 아니라 Statement로 취급된다.
	// 포인트로 집은 건 i++는 가능하지만 i++, j++로 하고 싶다면 이는 불가능하며, i, j = i+1, j+1 로 써야 한다는 것.
	// for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
	// 	fmt.Println(i, j)
	// }

	// 반복문 내에 3개의 statement를 넣을 필요는 없다.
	// i := 0
	// // 초기화 조건이 없지만 잘 작동함
	// for ; i < 5 ; i++ {
	// 	fmt.Println(i)
	// }
	// 작동 안함 : Initializer 자리가 정의되지 않았으니까?
	// for i < 5 ; i++ {
	// 	fmt.Println(i)
	// }

	// 반복문 내에 넣는 것과 밖에 넣는 것의 차이
	// 밖에 넣는 건 main() 내에서 정의된 변수임. 다른 곳에서도 쓸 수 있음
	// 안에 넣는 건 for 문 내에서만 정의된 변수임. i를 global하게 쓸 수 없음

	// Increment value도 없애보자 - 무한반복하겠죠? i는 종결 조건에 영원히 도달할 수 없으니까
	// i := 0
	// for i < 5 { // 실행조건문만 넣어줘도 잘 작동함. 양쪽에 ; 없어도 되네!
	// 	fmt.Println(i)
	// 	i++ // 실행문 내에 증가문을 넣어도 잘 작동한다
	// }

	// Undetermined Loop : 몇번 돌릴지 모르지만 적절한 타이밍에 빠져나가는 루프
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i == 5 {
	// 		break // 반복문 탈출하는 문구. 새삼스럽지도 않죠?
	// 	}
	// }

	// continue : 아래 스크립트 실행하지 말고 반복문 계속 돌려라
	// for j := 0; j < 10; j++ {
	// 	if j%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(j)
	// }

	// nested loop : 다중반복문
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			// (Loop 지정하기 전)
			// 1*3 출력되고 멈추겠지 히히 -> 실제로는 1 2 3 2 4 3 이 출력되고 멈춘다.
			// 왜냐하면 break은 가장 가까운 반복문만 break하기 때문이다. 즉 j 반복문이 break됨 (i=1에 대한)
			// 근데 i 반복문은 계속 실행되고 있잖아? 그래서 다시 (i=2에 대한) j 반복문이 호출됨
			// 그러면 같은 조건을 밖에도 붙여서 break하냐? 그렇게 안해도 된다.

			// 이중반복문 바깥에 Loop를 지정하고, break 뒤에 Loop를 넣으면 됨.
			if i*j >= 3 {
				break Loop
			}
		}
	}

	// Collections in for loop
	s := []int{1, 2, 3}
	fmt.Println(s)

	// 모든 item을 보여주고 싶다면
	for _, v := range s { // range를 쓰는 유일한 경우임. 반복 가능한 자료형 내부의 key와 value 값을 하나씩 보고 싶을 때 쓴다
		// key만 쓰고 싶은 경우는 k만 써서 추출할 수 있다. 근데 value만 쓰고 싶다면 go는 모든 변수가 쓰여야 한다는 조건이 있잖아?
		// 만약 value만 쓰고 싶다면 _를 활용하면 될 것 같다
		fmt.Println(v)
	}
}
