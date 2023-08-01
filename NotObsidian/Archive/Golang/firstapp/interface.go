package main

// 1. 인터페이스 정의
// type Shape interface {
// 	area() float64
// 	perimeter() float64
// }

// // 2. 인터페이스 구현 - area, perimeter의 메서드 2개를 구현하면 된다.
// // 근데 이걸 사용할 struct부터 정의해야겠져?
// type Rect struct {
// 	width, height float64
// }

// type Circle struct {
// 	radius float64
// }

// // For Rect
// func (r Rect) area() float64      { return r.width * r.height }
// func (r Rect) perimeter() float64 { return 2 * r.width * r.height }
// func (r Rect) custom() float64 {return 3 * r.width * r.height}
// // For Circle
// func (c Circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c Circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// 4. Empty Interface : {} 내부가 비어있다
// func Marshal(v interface{}) ([]byte, error)
// func Println(a ...interface{}) (n int, err error)

// 왜 쓸까? : 메서드를 전혀 갖지 않는 빈 인터페이스. Go의 모든 타입읜 최소 0개의 메서드를 구현한다
// 따라서 Go에서 모든 Type을 나타내기 위해 빈 인터페이스를 사용한다.
// 즉 어떠한 타입이라도 담을 수 있는 컨테이너 타입이다.
// 다른 언어에선 흔히 Dynamic Type이라고도 한다(cpp에서 void라고도 함)

// func main() {
// 	var x interface{}
// 	x = 1
// 	printlt(x)
// 	x = "Tom"

// 	printlt(x)

// 	var y int = 1
// 	printlt(y)
// }

// func printlt(v interface{}) {
// 	fmt.Println(v) //
// }

// 5. Type Assertion : x.(Type) 식으로 표현되며, 이는 x가 nil이 아니고, Type에 속한다는 점을 확인하는 표현이다.
// 만약 x가 nil 혹은 Type에 속하지 않는다면 runtime error이 발생한다.
func main() {
	var a interface{} = 1

	i := a // 별도로 지정되지 않는다면 interface{}는 dynamic type이다. 포인터 주소가 출력됨.
	j := a.(int)

	println(a) // 포인터 주소 출력
	println(i) // 포인터 주소 출력
	println(j) // 값 출력
}

// func main() {
// 	// 3. 인터페이스 사용 : 함수가 파라미터로 인터페이스를 받아들이는 경우
// 	r := Rect{10., 20.}
// 	c := Circle{10}

// 	showArea(r, c)
// }

// func showArea(shapes ...Shape) { // Shape가 인터페이스임
// 	// Shape 인터페이스를 구현한 타입 객체들을 파라미터로 받는다.
// 	// OOP를 따르지 않지만 OOP를 구현은 해놓은 모습인데.. 이 지점이 혼란스럽다.
// 	for _, s := range shapes {
// 		a := s.area() // 인터페이스에 있는 area라는 메서드 호출
// 		println(a)
// 	}
// }
