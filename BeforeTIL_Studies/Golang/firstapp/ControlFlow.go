package main

import (
	"fmt"
	"log"
)

func main() {
	// 1. Controlflow
	// 보통의 작업 과정은 위부타 아래로 진행됨
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	// 근데 defer를 넣으면 가장 뒤로 밀려남
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	// 작업 순서 : 위부터 아래 / 근데 defer가 있다는 걸 인식하고 지나감 -> main 함수가 "종료"된 이후 defer를 찾아서 실행함
	// 즉, defer는 main함수의 "끝"으로 보내는 게 아니라 "끝난 다음"으로 보냄

	// 근데 defer는 이런 특징도 있음 : LIFO로 실행됨
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	// defer 사용 예제
	// robots.txt를 가져오는 예제임
	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // defer는 이 close를 다루는 데 도움을 준다
	// // 왜냐! 많은 어플리케이션에서 body를 연 상태에서 작업할 필요성이 있음
	// defer res.Body.Close() // Close를 한 채로 작업한다면 어플에 버그가 생길 가능성이 있으므로, defer를 넣어준다

	// // 그래서 defer는 파일을 얻어오는 코드와 닫는 코드를 같이 배치할 수 있게 됨
	// // 순차적으로 Close()는 파일에서 긁는 과정이 끝난 뒤에 작성되어야 함. 한편 의미적으로는 파일을 열고 닫는 것에 관계가 있음
	// // 이를 defer를 이용해 의미가 있는 코드를 동시에 배치하면서도 실행되는 시간엔 차이를 둘 수 있다는 거임

	// // 그러나 강사는 다른 경고도 함
	// // loop 등을 통해서 많은 resource를 가져오고 request를 많이 해야 한다면, defer를 쓰는 건 좋은 방법이 아님
	// // main()이 끝난 뒤에 실행되므로, main()이 끝나기 전에 굉장히 많은 파일이 열려있다는 거고 이들을 main()이 끝남과 동시에 한꺼번에 실행한다는 거임
	// // 따라서 loop를 쓴다면 defer를 쓰지 말라는 것

	// robots, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// 예제 3
	// a := "start"
	// defer fmt.Println(a)
	// a = "end"
	// 예상 : end가 출력되지 않겠음? 가장 마지막으로 a가 갖는 값은 end니까
	// 실제 : 응 ~ start야~
	// 왜냐? : defer 뒤의 함수가 선언된 시점의 인수(argument)를 받음.
	// defer는 함수가 실행될 때 호출되는 게 아님! 순차적으로 내려가는 과정에서 호출됨
	// 따라서 이후에 a값이 바뀌더라도 defer 내부의 인수에 영향이 가지는 않음

	// 2. Panicking : Go 어플이 실행될 수 없고 이게 예외적인 상황인 경우
	// 그러나 예외적(Exceptional)이라는 표현을 쓰는 대신 Panic이라고 표현함 - 실행 불가능하고 뭘 해야 할지 몰라서 Panic이라는 표현을 쓴다
	// a, b := 1, 0
	// ans := a / b     // 1 / 0
	// fmt.Println(ans) // panic: runtime error: integer divide by zero

	// panic 예제 2.
	// fmt.Println("Start")
	// panic("Sth bad happened") // panic: Sth bad happened
	// fmt.Println("end")        // 실행되지 않음

	// panic 예제 3.
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!")) // response Writer
	// })

	// // 아래 panic 오류를 발생시키는 방법 : 다른 터미널에서 이 go 파일을 실행시킨다(이 파일을 실행시킨 채로)
	// // panic: (func() string) 0xc000004198
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error) // 만약 어떤 이유로 정상적으로 실행되지 않았다면 panic을 이용해서 오류가 발생했음을 유저에게 알려야 함
	// } // 실행 후 localhost:8080에 접속하면 "Hello Go!" 가 뜬다.

	// 예제 4. Panic이 발생했다면, 우리가 할 건 Recover이다
	// panic은 치명적일 필요가 없음
	fmt.Println("Start")
	// defer fmt.Println("This was Deferred") // 이거 실행되네 심지어 2번째로
	defer func() { // Anonymous Function
		if err := recover(); err != nil { // recover() : panick이 아니면 nil 반환 / panick이면 nil이 아님
			log.Println("Error:", err)
		}
	}() // 여기 () 부분이 Function Call 부분임
	panic("Something bad happened")
	fmt.Println("end")
	// 실행 순서 : main - defer - panic - handle return value

	// 왜 이게 중요하냐?
	// 1. Panic이 발생하더라도 defer function은 정상적으로 작동한다는 것
	// 2. 다른 함수에서 panic이 정의되고 그게 main 함수에서 작동했다고 치자
	// 그러면 panic 때문에 작동되지 않는 함수는 다른 함수 내부에만 적용되고, main 함수의 이후 코드들은 잘 작동한다
	// 즉 Panic이 발생했다면, 그 함수의 이후 코드들은 더 이상 신뢰할 수 없기 때문에 실행되지 않는다. (Recover은 된다)
	//  그러나 panic(err) 식으로 main()함수에 던져줄 수 있는데, 이 경우는 main()함수도 작동하지 않음(~_~ 모르겠다)

}
