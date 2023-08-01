package main

import (
	"fmt"
)

func main() {
	예제 1.
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	// if ~ ; : Initializer - pop, ok가 어떤 변수인지 알려주고, 중괄호 앞은 T/F가 와야 하므로 ok가 와있음
	if pop, ok := statePopulations["oho"]; ok{ // ok는 True / False 역할임(보통의 조건문 역할)
		fmt.Println(pop)
	}

	// fmt.Println(pop) // pop은 조건문 내에서 만들어진 변수이므로 global하지 않음
	

	// 예제 2.
	// number := 50
	// guess := 105
	// if guess < 1 || guess > 100 { // 2가지 테스트를 수행함, or 이 ||로 쓰였다
	// 	fmt.Println("The guess must be between 1 and 100!")
	// }
	// if guess > 1 && guess <= 100 { // AND가 &&로 쓰였다
	// 	if guess < number { // less than
	// 		fmt.Println("Too low")
	// 	}
	// 	if guess > number { // greater than
	// 		fmt.Println("Too high")
	// 	}
	// 	if guess == number { // equal
	// 		fmt.Println("You Got it !")
	// 	}
	// 	// guess에 따른 실행 결과 : 30 - Too low / 70 - Too High / 50 - You Got it !
	// 	fmt.Println(number <= guess, number >= guess, number != guess) // 3가지 조건문 모두 Boolean을 반환
	// }
	// fmt.Println(!true) // NOT Operator
	// fmt.Println(!false)

	// 예제 3. Short Circuiting - returnTrue() 함수 추가
	number := 50
	guess := 105 // -5를 집어넣었을 때 returnTrue()에서 print해야 할 "Returning True"가 작동되지 않음
	// 이를 Short Circuiting이라고 함
	// OR test의 하나의 조건이 True를 만족한다면, 그 뒤의 조건들은 더 이상 수행되지 않음
	// guess := 105여도 마찬가지. returnTrue()가 true이므로 guess>100의 조건문은 수행되지 않는다.

	if guess < 1 || returnTrue() || guess > 100 { // 3가지 테스트 동시 수행 : 가운데 True가 있으니 무조건 True 조건 만족함
		fmt.Println("The guess must be between 1 and 100!")
	}

	// ShortCircuiting은 AND 조건문에도 작동한다 : 하나의 조건이라도 False라면 뒤의 조건문은 수행되지 않는다
	if guess > 1 && guess <= 100 { // 한편 위의 조건과 병렬적으로 세워지므로 이것도 작동함
		if guess < number { // less than
			fmt.Println("Too low")
		}
		if guess > number { // greater than
			fmt.Println("Too high")
		}
		if guess == number { // equal
			fmt.Println("You Got it !")
		}
	}

	// 예제 4. 위의 조건문이 지저분해요!
	number := 50
	guess := -5
	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess < 1 || guess > 100 { // if - else if - else는 선형적이므로 같은 조건문이 True인 조건이 나오면 멈춤
		fmt.Println("The guess must be less than 100!")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You Got it !")
		}

		fmt.Println(number <= guess, number >= guess, number != guess)
	}

	// 예제 5. 수치형 데이터를 비교할 때..
	// myNum := 0.123
	// if myNum == math.Pow(math.Sqrt(myNum), 2) { // 제곱근 씌우고 다시 제곱하는 거 같다.
	// 	fmt.Println("These are the same")
	// } else {
	// 	fmt.Println("These are different")
	// }
	// // 0.1일 땐 True인데 0.123에선 False가 나옴 : 정확한 값을 비교하지 않기 때문임(자르는 경우가 있으니까)

	// // 따라서 위의 조건문은 이렇게 돌리는 게 바람직함 : Threshold를 정해주는 것
	myNum2 := 0.123
	if math.Abs(myNum2/math.Pow(math.Sqrt(myNum2), 2)-1) < 0.001 { // 즉 이 2개를 처리한 값이 1%보다 작냐
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}

	// 예제 6. Switch

	// if의 특별한 구문 / 실제로 사용할 때 비슷한 비교를 많이 하므로 switch를 쓰는 게 더 유용하다
	switch 2 { // switch 뒤에 오는 2를 tag라고 함
	// case 뒤의 값을 tag와 비교함 - true인 것만 실행됨
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	// true가 없다면 default가 실행됨
	default:
		fmt.Println("Not one or two")
	}

	// case 뒤에는 여러 개의 숫자가 올 수 있다
	// switch 4 {
	// case 1, 5, 10:
	// 	fmt.Println("One, five or ten")
	// case 2, 4, 6:
	// 	fmt.Println("two, four or six")
	// default:
	// 	fmt.Println("Another Number")
	// }

	// case 뒤에 오는 값은 고유해야 함. 겹치면 switch 문 실행 안됨

	// switch i := 2 + 3; i { // ; 뒤의 i가 테스트 값임
	// case 1, 5, 10:
	// 	fmt.Println("one, five or ten")
	// case 2, 4, 6:
	// 	fmt.Println("two, four or six")
	// default:
	// 	fmt.Println("Another Number")
	// }

	// // Tagless도 있다 : 조건이 겹침에도 작동하는 것에 주목(1번쨰 case만 작동하지만)
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // 의미 : 다음 케이스의 실행을 의도적으로 원한다
	// case i <= 20:
	case i >= 20: // 그러나 fallthrough 를 넣으면 2번째 case가 조건 만족하지 않음에도 작동한다.
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Type Switch
	var i interface{} = [3]int{} // tag i에 interface{}가 붙음 : 어떤 데이터 타입이든 올 수 있다. collection 타입까지도.
	switch i.(type) {            // i.(type)은 이외에도 많이 쓰이는 문법이니 알아두면 좋다. type을 반환함.
	case int:
		fmt.Println("i is an int")
		break // break도 잘 작동함
		fmt.Println("This will print too")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	case [2]int:
		fmt.Println("i is [2]int") // Array의 경우 크기까지 따짐
	default:
		fmt.Println("i is another type")
	}
}

// func returnTrue() bool {
// 	fmt.Println("Returning true")
// 	return true
// }
