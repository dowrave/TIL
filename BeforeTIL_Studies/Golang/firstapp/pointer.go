package main

import (
	"fmt"
)

func main() {
	// 1. 포인터
	var a int = 42
	// b := a
	var b *int = &a     // 데이터타입 앞에서 포인터 선언(변수명이 아니라), 그 주소는 a의 주소이다
	fmt.Println(a, b)   // 42, 0xc0000aa058
	fmt.Println(&a, *b) // 0xc0000aa058, 42

	/* 나중에 헷갈리니까 확실히 정리하고 가자
	b라는 변수는 a의 메모리 주소를 저장한다. b라는 변수는 a의 값을 저장하고 있지 않다.
	b는 a의 메모리 주소를 저장함
	이 때, b의 값은 *b로 호출할 수 있고, a의 메모리 주소는 a로 호출할 수 있음
	(기본적으로 a는 값을, b는 주소를 호출한다) */

	// 여기서 값을 변경해보자
	// a = 27
	// fmt.Println(a, *b) // 둘 모두가 같은 주소를 지정하므로..
	// *b = 14
	// fmt.Println(a, *b) // 이것도 마찬가지

	// 2. Pointer Arithmetic(포인터 산술) - Go는 Simplicity가 핵심이라서, Go에서는 없음
	a := [3]int{1, 2, 3}
	b := &a[0] // b는 0번째 element의 주소
	c := &a[1] // c도 1번째 element의 주소
	// go는 포인터에 연산을 허용하지 않는다
	// c := &a[1] - 2 // C나 CPP에선 가능한 연산임
	fmt.Printf("%v %p %p \n", a, b, c) // 0xc000014138 0xc000014140 : 내꺼에선 각 int가 2bytes 차지하나보다

	// 만약 Pointer Arithmetic을 사용하고 싶다면, unsafe package를 참조할 것

	// 3. *, & 없이도 이 값은 성립함. 단 출력되는 값이 다름. 이 방법의 장점은 나중에 설명함
	// var ms *myStruct
	// // ms = &myStruct{foo: 42}
	// // new라는 함수를 쓸 수도 있다 : 단 바로 값을 넣을 수는 없다
	// ms = new(myStruct) // new로 선언된 myStruct는 0이라는 값을 저장함
	// fmt.Println(ms)    // &{42} : ms는 기본적으로 주소를 저장하고 있으며, 값 field에 42를 저장함

	// 초기화값에 대해서
	var ms *myStruct
	fmt.Println(ms) // <nil> : 선언된 시점에선 nil이라는 값을 저장함
	// 이대로 진행되면 crash가 발생하므로 처리하고 갈 것
	ms = new(myStruct)
	fmt.Println(ms) // &{0} : new로 선언된 지점에선 0 값을 저장함

	// 이런 식으로 처리할 수 있다
	var ms *myStruct // 여기서 ms는 주소(=포인터)임
	// fmt.Println(ms)        // nil : 포인터의 Zero Value임.
	ms = new(myStruct)     // dereference : 주소를 통해 값에 접근(역참조) . 이게 없으면 nil 값이 있는 주소를 역참조했다는 오류가 뜸
	(*ms).foo = 42         // *ms.foo가 안됨
	fmt.Println((*ms).foo) // 포인터 내부의 속성에 접근하려면 (*ms)로 감싸야한다는 거넹
	// 근데 사실 이거도 된다 (다른 언어에선 안된다)
	ms.foo = 42
	fmt.Println(ms.foo)
	/* 이 기능은 포인터를 사용하는 다른 언어에선 작동하지 않음 : ms 자체는 포인터고, 포인터 자체에는 foo가 들어가는 field가 없기 때문이다.
	이건 컴파일러가 ms라는 포인터의 오브젝트에 접근한다는 걸 안다는 의미가 됨 */

	// 4.
	a := []int{1, 2, 3}
	b := a // a가 array이든 slice든 map이든 복사함. 근데 array는 자체가 값을 갖고 있는 반면
	// slice, map은 주소인 포인터만 가짐 -> 따라서 복사된 b는 a의 값들을 참조하게 됨
	// slice는 자체가 데이터를 저장하는 게 아님 : slice는 1번째 원소 위치를 지정하는 포인터를 저장함
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)

	// 따라서 어플을 만들 때 slice나 map의 사용에도 유의해야 함
	// 복사했다고 생각하고 값을 바꿨는데 원본 데이터의 값이 바뀌기 때문.

	// Primitives, array, struct는 복사되었을 때 "포인터를 사용하지 않는다면" 전체 구조가 복사된다.
}

type myStruct struct {
	foo int
}
