package main

// 8. closure
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 1. Pass by Value
	// msg := "Hello"
	// say(msg)

	// 2. Pass by Reference
	// say_ref(&msg) // Hello : msg의 주소 값이 들어가므로
	// println(msg) // say_ref에서 msg 주소에 있는 값이 change로 바뀌었음

	// 3. Variadic Function(가변인자함수)
	// say("This", "is", "a", "book")
	// say("Hi")

	// 4. 함수 리턴값
	// count, total := sum(1, 7, 3, 5, 9)
	// println(count, total)

	// 5. 익명 함수(Anonymous Function)
	// sum := func(n ...int) int {
	// 	s := 0
	// 	for _, i := range n {
	// 		s += i
	// 	}
	// 	return s
	// }
	// result := sum(1, 2, 3, 4, 5)
	// println(result)

	// 6. 일급함수
	// Go의 함수는 다른 함수의 파라미터로 전달하거나 리턴값으로 사용될 수 있다
	// 즉 함수의 입력 파라미터나 리턴 파라미터로서 함수가 사용될 수 있다.
	// 이를 위해서는 익명 함수를 변수에 할당한 후 변수로 전달하거나, 직접 다른 함수 호출 파라미텅 ㅔ함수를 적는 방법이 있다

	// add := func(i int, j int) int { // 익명함수 정의
	// 	return i + j
	// }

	// r1 := calc(add, 10, 20) // 익명함수를 calc라는 함수에 넣음. calc는 func와 int형 2개를 input으로 받음
	// println(r1)

	// r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	// println(r2)

	// 8. closure
	next := nextValue()

	println(next())
	println(next())
	println(next())

	anotherNext := nextValue()
	println(anotherNext())
	println(anotherNext())
}

// func say(msg string) {
// 	println(msg)
// }

// func say_ref(msg *string) {
// 	println(msg)
// 	println(*msg)
// 	*msg = "Changed"
// }

// 3. Variadic Function
// func say(msg ...string) {
// 	for _, s := range msg {
// 		println(s)
// 	}
// }

// 4. 함수 리턴값
// C는 함수가 void 혹은 1개만 리턴할 수 있지만
// Go는 0개 ~ 복수개 까지 가능하다
// 또한 Named Return Parameter라는 기능도 제공된다.

// func sum(nums ...int) (int, int) { // return parameter가 여러개라면 여기서 자료유형을 지정해준다
// 	s := 0
// 	count := 0
// 	for _, n := range nums { // range는 반복가능한 객체의 idx, value를 반환함
// 		s += n
// 		count++
// 	}
// 	return count, s // func 문장의 (int, int) 타입을 따라감

// 위 함수는 Named Return Parameter 기능을 이용하면 다음과 같다
// func sum(nums ...int) (count int, total int) {
// 	for _, n := range nums {
// 		total += n
// 	}
// 	count = len(nums)
// 	return // return 값을 함수 선언되는 문장으로 올린다. 그러나 return 값은 꼭 써줘야 한다.
// }

// 6. 일급함수
// func calc(f func(int, int) int, a int, b int) int { // 파라미터로 함수를 선언
// 	// f func(int, int) int : input으로 int 형 2개를 받고 1개의 int return
// 	// calc는 3개의 input을 받고(func, a, b) 1개의 int return
// 	result := f(a, b)
// 	return result
// }

// 7. Type으로 함수 원형 정의하기
// type : 구조체, 인터페이스 등 커스텀 타입을 정의하기 위해 사용되는데, 함수 원형을 정의할 때 쓰이기도 한다.
// 위 예제에서 func(x int, y int) int 함수 원형이 코드 상에 계속 반복되는데, 이 때 type문을 정의하면 원형을 간단히 표현할 수 있음
// type calculator func(int, int) int

// func calc(f calculator, a int, b int) int {
// 	result := f(a, b)
// 	return result
// }
