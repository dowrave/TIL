package main

import "fmt"

type Rect struct { // 구조체는 type으로 정의한다 복습
	width, height int
}

func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

// 포인터 리시버를 이용
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	fmt.Println("ValueReceiver : ", rect.width, area)
	area2 := rect.area2() // 메서드 호출
	fmt.Println("PointerReceiver : ", rect.width, area2)
	area3 := rect.area()
	fmt.Println("ValueReceiver 2 : ", rect.width, area3)
}
