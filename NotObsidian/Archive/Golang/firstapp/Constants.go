package main

import (
	"fmt"
)

// 3. Shadowing : 패키지 수준에서 만들기
// const a int16 = 27 // 의외로 const는 불변은 아니다. 패키지 수준에서 정의된 걸 함수 레벨에서 재정의할 수 있음
// 그러나 어쩄든 값이 변하는 상황이 생기므로 저런 식으로 이중 정의를 하는 것은 권장하지 않는다

// 5. iota : enumerated const
// const a = iota // iota : 열거 상수를 만들 때 사용

// 위보다는 아래처럼 씀
// const (
// 	a = iota // 0, int
// 	b = iota // 1, int
// 	c = iota // 2, int
// )

// 혹은 이런 방식 : 하나만 지정하면 나머지는 순차적으로 올라감
// const (
// 	a = iota
// 	b // = iota
// 	c // = iota
// )

// // 새로운 블럭에 iota를 지정하면 다시 0부터 시작함
// const (
// 	a2 = iota
// )

// 그러면 iota를 어떻게 이용할 수 있을까?
// const (
// 	// errorSpecialist = iota
// 	_ = iota
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )

// iota 활용 예제 : bit shifting
// 여러 변수를 이런 식으로 지정할 때 매우 좋긴 하겠다
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // 1 << 20
	GB                    // 1 << 30
	TB
	PB
	EB
	ZB
	YB
)

// iota 활용 예제 2
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAFrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// 1. 상수의 이름을 지정하는 방법

	// const myConst // 함수 내부에서만 쓰겠다
	// const MyConst // Export 할 용도로 변수를 정할 때 쓴다

	// 2. const로 지정될 수 있는 것들
	// const myConst float64 = math.Sin(1.57) /* 실행 불가능 : math.Sin(1.57) 값은 execute 시에 정해짐. const는 이를 못 받쳐줌 */
	// myConst = 27 상수로 지정했기 때문에 값 변경 x
	// fmt.Printf("%v, %T\n", myConst, myConst)

	// 그 외에 기본적은 int, string, float, boolean 타입은 잘 받음
	// fmt.Printf("%v, %T\n", a, a)

	// const a int = 42
	// var b int16 = 27
	// const b string = "foo"
	// const c float32 = 3.14
	// const d bool = true
	// fmt.Printf("%v, %T\n", a+b, a+b) // 상수와 변수의 연산 : 가능 : 데이터 유형은 다르니까
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)
	// fmt.Printf("%v\n", d)

	// 4. 컴파일러를 이용하여 유형 추론하기
	// const c = 42 // 위 예제와 차이라면 데이터 유형(int)을 지정하지 않았음 / d가 없다면 기본 int로 지정
	// var d int16 = 27
	// fmt.Printf("%v, %T\n", c, c)
	// fmt.Printf("%v, %T\n", c+d, c+d)
	// 왜 이게 되느냐? 컴파일러가 c = 42를 볼 때 "이건 바뀌지 않아야 할 42, int"라고 보는 게 아니라
	// c라는 기호를 쓰는 모든 상황을 보고, 그때그때 c의 값을 넣어주기 때문임

	// 5. Enumerated Constant
	// fmt.Printf("%v, %T\n", a, a)
	// fmt.Printf("%v, %T\n", b, b)
	// fmt.Printf("%v, %T\n", c, c)
	// fmt.Printf("%v, %T\n", a2, a2)
	// fmt.Printf("%v, %T\n", d, d)

	// var specialistType int = catSpecialist
	// fmt.Printf("%v\n", specialistType == catSpecialist)

	// iota는 기본적으로 0, 1, 2, ... 순으로 올라감
	// 그렇다면 다음 예제를 봅시다
	// var specialistType int // 초기화 X
	// fmt.Printf("%v\n", specialistType == catSpecialist)

	// 초기화를 하지 않았는데도 실행하면 True임 : 왜냐면 변수 초기화값과 iota 초기화 값이 0으로 동일하기 때문임
	// 이를 방지하는 방법

	// 1. catSpecialist 위에 errorSpecialist 값을 넣고 여기서부터 출발하게 만듦
	// 2. errorSpecialist 대신 _ 를 집어넣는 것 : error를 굳이 확인할 필요가 없고, 1부터 출발시키고 싶을 경우
	// _를 쓰는 건 메모리를 아끼는 측면에서도 좋음 : 할당하는 주소가 없기 때문

	// iota 활용 예제 : bitshifting
	fileSize := 4000000000.
	fmt.Printf("%.2fGB",
		fileSize/GB) // GB값이 어떻게 지정되었느냐 보면
	1 << (10 * iota)
	10 * iota = 30
	// 즉 왼쪽으로 30자리 밀라는 뜻이잖아? 그럼 2^30번째 자리에 1이 있는 거임

	// iota 활용 예제 2
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	// 1 / 100 / 100000 들을 or을 돌리면?
	fmt.Printf("%b\n", // 비트 출력
		roles) // 100101

	// 즉 찾아낸 결과가 뭐다? : Admin이면서, Financials을 볼 수 있으면서, Europe을 볼 수 있는 유저다
	// 이런 식으로 어떤 것에 접근가능한 사람인지 확인할 수 있음. ID까지는 못하더라도.
	fmt.Printf("Is Admin? %v", isAdmin&
		roles == isAdmin) // 이거는 boolean 질문. true / false 반환
	fmt.Printf("Is HQ? %v", isHeadquarters&
		roles == isHeadquarters)
}
