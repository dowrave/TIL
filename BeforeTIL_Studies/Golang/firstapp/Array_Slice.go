package main

import (
	"fmt"
)

func main() {

	// 이런 식으로 자료를 다루는 건 비효율적일 것임
	// grade의 수가 바뀔 수도 있고..
	// grade1 := 97
	// grade2 := 85
	// grade3 := 93
	// fmt.Printf("Grades: %v, %v, %v", grade1,
	// 	grade2, grade3)

	// 그래서 이런 식으로 해라 : ARRAY
	// [들어갈 갯수]자료형{값1, 값2, ...}
	// grades := [3]int{97, 85, 93}
	// fmt.Printf("Grades: %v", grades)
	// Array의 장점
	// 1. 데이터를 다루기 매우 유용함
	// 2. 메모리에 남아 있음 : 액세스가 매우 빠름

	// 근데 들어갈 갯수 하나하나 정하면 까다롭잖음? 그래서 ...로 넣어줘도 됨
	// grades := [...]int{97, 85, 93} // 통과시켜주는 값만큼 array 공간을 만듦
	// fmt.Printf("Grades: %v", grades)

	// 혹은 구조는 정하는데 값은 못정했다고 하면
	// var students [3]string
	// fmt.Printf("Students : %v\n", students)
	// students[0] = "Lisa"
	// fmt.Printf("Students : %v", students)

	// Array 특 : 포인터는 array의 가장 1번째 장소를 지정함
	// var students [5]string
	// fmt.Printf("Students : %v\n", students)
	// 인덱스 지정 순서는 상관 없음
	// students[0] = "Lisa"
	// students[1] = "Ahmed"
	// students[2] = "Arnold"
	// fmt.Printf("Students : %v\n", students)
	// 인덱싱 생략함. students[1]로 호출하면 %v값에 나오는 거

	// Built-in Function
	// len()
	// fmt.Printf("Number of Studetns : %v\n", len(students))

	// Array in Array
	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1, 0, 0}
	// identityMatrix[1] = [3]int{0, 1, 0}
	// identityMatrix[2] = [3]int{0, 0, 1}

	// fmt.Println(identityMatrix)

	// 다른 언어와 Golnag이 Array에서 다른 점
	// a := [...]int{1, 2, 3}
	// b := a // Literal Copy
	// b[1] = 5
	// fmt.Println(a) // b:=a 인데 같은 장소를 지정하지 않음
	// 즉 b는 a를 카피한 뒤 다른 장소에 저장된 array임
	// fmt.Println(b)

	// 즉 array가 클수록 := 를 써서 복사하는 방식은 프로그램을 다운 시킬 수 있다
	// 그렇다면 이를 방지하는 방법은?
	// "포인터" - 장소를 지정하는 변수는 copy를 방지할 수 있다
	// c := &a // cpp에서의 사용과 동일함
	// c[1] = 5
	// fmt.Println(a)
	// fmt.Println(c)

	// Slice : slice라는 데이터 유형이 따로 있는 듯? array 같이
	// a := []int{1, 2, 3} // slice 선언 . []타입{값}
	// fmt.Println(a)
	// fmt.Printf("Length : %v\n", len(a)) // len도 3

	// // capacity라는 builtinfunction을 써보자
	// fmt.Printf("Capacity : %v\n", cap(a))
	// b := &a
	// b[1] = 5
	// 위 두 줄로 아래 값들이 바뀌는 건 당연하잖아? 스킵

	// Slice 예제
	// arr_a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := arr_a // copy : b는 새로운 주소에 할당됨
	// b[5] = 12
	// fmt.Printf("%v\n", arr_a)
	// fmt.Printf("%v\n", b)

	// sli_a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// c := sli_a // refer : 원본 주소의 데이터 변경
	// sli_a[5] = 58
	// fmt.Printf("%v\n", sli_a)
	// fmt.Printf("%v\n", c)

	// slice를 만드는 다른 방법 : make
	// a := make([]int, 3) // slice의 length (len()으로 호출됨)
	// fmt.Println(a)
	// fmt.Printf("Length : %v\n", len(a))
	// fmt.Printf("Capacity : %v\n", cap(a))

	// // make에는 3번째 argument를 만들 수도 있다 : capacity
	// b := make([]int, 3, 100)
	// fmt.Println(b)
	// fmt.Printf("Length : %v\n", len(b))
	// fmt.Printf("Capacity : %v\n", cap(b))

	// len과 cap이 뭔지에 대한 예제
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := a[3:5]
	c := a[1:3]                 // 기대 : len 2 cap 6
	fmt.Println(len(b), cap(b)) // 2, 4
	fmt.Println(len(c), cap(c))

	// append를 통한 cap 관찰
	s := []int{}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, %d, %p\n", len(s), cap(s), s)
		s = append(s, i)
	}
	// c := []int{}
	// fmt.Println(c)
	// fmt.Printf("Length : %v\n", len(c))
	// fmt.Printf("Capacity : %v\n", cap(c))
	// c = append(c, 1)
	// fmt.Println(c)
	// fmt.Printf("Length : %v\n", len(c))
	// fmt.Printf("Capacity : %v\n", cap(c))
	// c = append(c, 2, 3, 4, 5)
	// fmt.Println(c)
	// fmt.Printf("Length : %v\n", len(c))
	// fmt.Printf("Capacity : %v\n", cap(c))

	// c = append(c, []int{6, 7, 8, 9}) // 2번쨰 argument는 반드시 interger가 와야 함
	// 따라서 위 방식으로는 사용할 수 없다

	// 대신 JS에서 Spread dot이라고 불리는 ...에 대해서는,
	// ...는 slice 내부의 원소들을 퍼뜨리는 역할을 함
	// c = append(c, []int{6, 7, 8, 9}...)
	// // c = append(c, 6, 7, 8, 9) 와 같음
	// fmt.Println(c)
	// fmt.Printf("Length : %v\n", len(c))
	// fmt.Printf("Capacity : %v\n", cap(c))

	// Stack Operation
	// a := []int{1, 2, 3, 4, 5}
	// b := a[:len(a)-1]
	// fmt.Println(b)

	// a := []int{1, 2, 3, 4, 5}
	// fmt.Println(a)
	// // 중간의 요소 하나만 제거
	// b := append(a[:2], a[3:]...)
	// fmt.Println(b)

	// fmt.Println(a) // [1 2 4 5 5]???
	// 이런 경우가 생기므로 중간의 요소를 제거하고 싶다면
	// 해당 슬라이스를 참조하는 다른 변수가 없어야 함

	// 이걸 다루려면 반복문을 한번 보고 와야 한다고 함
}
