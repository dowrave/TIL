package main

import "fmt"

func main() {
	// 1. Bool
	// var n bool = true
	// fmt.Printf("%v, %T\n", n, n)

	// Logical Test
	// n := 1 == 1
	// m := 1 == 2
	// fmt.Printf("%v, %T\n", n, n) // true, bool
	// fmt.Printf("%v, %T\n", m, m) // false, bool

	// 이건 어떨까
	// var n bool // 선언만 함
	// fmt.Printf("%v, %T", n, n)
	// 초기화 하지 않으면 False로 나가는 듯?

	// 2. Numerical
	// n := 42
	// fmt.Printf("%v, %T\n", n, n)

	// // Unsigned Integer
	// var n uint16 = 42
	// fmt.Printf("%v, %T\n", n, n)

	// // 연산 : +, -, *, /, %
	// // 정수 타입에 대해 진행되었다면 /는 몫만 나옴
	// // 다른 타입의 정수를 더하면 오류 발생함
	// var a int = 10
	// var b int8 = 3
	// fmt.Println(a + b) // int와 int8이 다르기 때문에 계산 못해요~
	// fmt.Println(a + int(b))

	// // 즉 Go는 암시적인 가정을 하지 않으므로, 하나하나 명시적으로 변환해줘야 함

	// // 비트 연산 : &, |, ^, &^
	// // 실제로는 32 ~ 64 비트지만 앞의 0들을 무시함
	// a := 10 // 1010
	// b := 3 // 0011
	// fmt.Println(a & b) // 0010 (and : 1&1만 1) = 2
	// fmt.Println(a | b) // 1011 (or : 0|0 뺴고 1) = 11
	// fmt.Println(a ^ b) // XOR : 두 비트가 다르면 1 : 1001 = 9
	// fmt.Println(a &^ b) // clear : b에서 1인 자리들을 a에서 0으로 처리함 : 1000 = 8

	// // 비트 시프팅
	// a := 8              // 1000
	// fmt.Println(a << 3) // 1000000 = 2^6 (왼쪽으로 3칸 이동)
	// fmt.Println(a >> 3) // 0001 = 2^0 (오른쪽으로 3칸 이동)

	// 3. Float
	// var n float32 = 3.14 // 이는 오류 발생 : 13.7e72는 범위 밖이기 때문
	// n := 3.14
	// fmt.Printf("%v, %T", n, n) // 기본적으로 float64로 생성됨
	// n = 13.7e72 // 작은 e든 큰 e든 같은 듯. 저장하니까 E가 e로 바뀜
	// n = 2.1e14
	// fmt.Printf("%v, %T", n, n)
	// var a float32 = 1.02
	// var b float32 = 4.79
	// fmt.Printf(a & b)
	// 마찬가지로 float64와 float32은 연산 불가능

	// 연산자도 동일함
	// 연산자 차이 : %, << , >>는 사용 불가능

	// 복소수 : complex
	// 예제 하나를 봅시다
	// var n complex128 = 1 + 2i // 64와 128이 있다
	// var n complex128 = complex(5, 12) // 실수분, 허수분
	// fmt.Printf("%v, %T\n", n, n)

	// 복소수의 연산도 동일함
	// 복소수를 분리하는 방법은? - builtin function을 이용한다
	// fmt.Printf("%v, %T\n", real(n), real(n))
	// fmt.Printf("%v, %T\n", imag(n), imag(n))
	// 단 위에서 complex 64 / 128을 썼다면 실수분 허수분은 각각 float32, 64로 나옴

	// 4. String
	// s := "This is a String"
	// fmt.Printf("%v, %T\n", s, s)

	// String은 Array처럼 취급할 수 있다
	// fmt.Printf("%v, %T\n", s[2], s[2]) // 105, uint8
	// 머선 일이고? string은 byte의 별칭이다. 따라서
	// fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// 그러나 이런 식으로 작동하지는 않는다.
	// s[2] = "u"
	// fmt.Printf("%v, %T\n", s, s)

	// 문제 1. 문자열을 바이트에 할당할 수 없다 -> 변환 필요
	// 문제 2. 문자열의 값을 조작할 수 없음

	// 1) 문자열 연결
	// s1 := "This is a String"
	// s2 := "That is also a String"
	// fmt.Printf("%v, %T\n", s1+s2, s1+s2)

	// 2) 바이트 조각으로 변환
	// b := []byte(s1) // s1의 각 문자를 ASCII로 변환해서 array에 담는 것 같다
	// fmt.Printf("%v, %T\n", b, b)

	// 바이트 슬라이싱이 중요함
	// 많은 함수는 바이트로 잘리는데, 이는 유연성에 중요한 역할
	// 바이트로 쪼개서 보내는 건 많이 쓰이므로 알아두면 좋다

	// 5. Rune
	// r := 'a'                     // 하나의 문자임에 주목
	// fmt.Printf("%v, %T\n", r, r) // 97, int32

	// var r2 rune = 'a'
	// fmt.Printf("%v, %T\n", r2, r2) // 97, int32

	// Rune이 뭔데 십덕아
	// UTF-32, 처리하기 위해선 별도의 특수한 메소드가 필요함
	// 예시) strings.Reader#ReadRune
	// UTF-8 인코딩을 쓰는 Golang은 int32 타입으로 표시된다

	// String과 Rune
	// String : bytes로 만들어짐, rune으로 표현될 수 있는 유효한 character 들을 포함할 수 있음
	// string을 rune(string)을 통해 rune의 array로 바꿀 수 있음
	// ASCII 문자에 대해서는 rune 값은 byte 값과 같음

	// 다양한 문자를 바이트로 표현하기 위해 rune이라는 타입을 도입했다고 이해해도 될까?

	a := "아버지가방에들어가신다"
	a_rune := []rune(a)

	a_byte := []byte(a)

	fmt.Printf("%v, %T\n", a_rune, a_rune)
	// [50500 48260 51648 44032 48169 50640 46308 50612 44032 49888 45796], []int32
	fmt.Printf("%v, %T\n", string(a_rune[3]), a_rune[3]) // 가, int32

	fmt.Printf("%v, %T\n", a_byte, a_byte)

	fmt.Printf("%v, %T\n", string(a_byte[3]), a_byte[3]) // ë, uint8

	/* [236 149 132 235 178 132 236 167 128 234 176 128 235 176 169 236 151 144 235 147 164
	236 150 180 234 176 128 236 139 160 235 139 164], []uint8 */

	// var b string = '가'
	// fmt.Printf("%v, %T\n", b, b) // cannot use '가' (untyped rune constant 44032) as string value in variable declaration)

	// 대충 드는 생각으로는, 저런 방법을 취한다면 글자를 슬라이싱할 때 Rune 방법이 훨씬 편함.
}
