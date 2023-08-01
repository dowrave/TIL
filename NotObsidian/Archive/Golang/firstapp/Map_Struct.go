package main

import (
	"fmt"
	"reflect"
)

// 2. Struct : 얘도 map과 마찬가지로 collection 타입이다.
// 여기서도 명명법은 동일함. Export할 건 앞글자를 대문자로,
// 함수 내에서만 쓴다면 소문자만 쓰고..
type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

// 3. Embedding - Go는 Inhertiance와 비슷한 Composition = Embedding이라는 모델을 지원함
type Animal struct {
	Name   string `required max:"100"` // tag
	Origin string
}

type Bird struct {
	Animal // Animal Animal로 넣지 않음 : 즉 Bird 내에 Animal이라는 Name Field를 두지 않았음
	// Animal이라는 위에서 선언한 구조체 자체만 아래 구조체에 집어넣음
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// 1. map이 뭐죠? : map[keytype]valuetype{} 파이썬으로 치면 dict 개념이지만 key&value의 type을 지정해줘야 함
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	} // 이 방법을 Literal Syntax라고 함

	// Equivalency Checking이 불가능한 자료 유형들 : slice, map, 몇몇 function
	m := map[[]int]string{}  // slice : 불가능
	m := map[[3]int]string{} // array는 가능

	// 1-1. make function으로도 만들 수 있음 : 변수를 넣을 시간은 없고 만들어두고 싶을 때
	// statePopulations := make(map[string]int)
	// statePopulations = map[string]int{
	// 	"California":   39250017,
	// 	"Texas":        27862596,
	// 	"Florida":      20612439,
	// 	"New York":     19745289,
	// 	"Pennsylvania": 12802503,
	// 	"Illinois":     12801539,
	// 	"Ohio":         11614373,
	// }

	// // map의 출력 순서 : 보장되지 않음(우리가 넣은 대로 나오지 않음)

	// // map에 key&value pair 추가하기
	statePopulations["Georgia"] = 10310371

	// //key & value pair 삭제하기
	delete(statePopulations, "Georgia")

	// fmt.Println(statePopulations)

	// //없는 key를 집어넣으면 0이 출력된다.
	fmt.Println(statePopulations["Georgia"])
	// /* 이건 문제가 될 수 있음 : value값이 0인지 없어서 0인지 구분되지 않기 때문임.
	// 따라서 이를 파악하기 위해 다음 방식을 따른다 : 강의에선 comma를 쓴다는 식으로 언급함*/
	_, ok := statePopulations["Oho"]
	fmt.Println(ok) // 0, False

	// // map 원소 갯수
	// fmt.Println(len(statePopulations)) // 7

	// 1-2. Map을 참조 = 그 값은 "주소를 참조함"
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	// 2. Struct
	// Struct의 장점 : 여러 자료 유형을 하나에 다룰 수 있다는 것
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	// 이름을 지정할 필요도 없음 - 문법적으로 유효하나, 유지보수 측면에서 좋은 방법이 아님 (추후 수정할 상황을 고려)
	// 예를 들어 구조체 내부 구조 선언 자체를 바꾼다고 생각해보자. 의미하는 바가 바뀔 수 있음
	// 이를 Positional Syntax라고 함
	// aDoctor := Doctor{
	// 	3,
	// 	"Jon Pertwee",
	// 	[]string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }
	// 자꾸 없어지냐
	// fmt.Println(aDoctor.actorName) // struct 내부 속성 호출은 .을 이용함

	// Struct를 선언하는 다른 방법 : Anonymous Struct(구조체 이름이 선언되지 않았으니까)
	// aDoctor := struct{ name string }{name: "John Pertwee"}
	// fmt.Println(aDoctor)
	// 오래 사용하지 않을 경우에만 쓰임

	// map, slice와 달리 struct의 참조는 독립적임(별도의 주소에 저장됨)
	// aDoctor := struct{ name string }{name: "John Pertwee"}
	// anotherDoctor := aDoctor
	// 따라서 동일한 주소를 참조하고 싶다면 포인터 &를 쓰자
	// anotherDoctor := &aDoctor
	// anotherDoctor.name = "Tom Baker"
	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)

	// 3. Embedding : Go는 전통적인 OOP를 지원하지 않는다
	// 예를 들어, Inheritance를 지원하지 않음 - 헐 그럼 어케 만들어요
	// 위의 Animal, Bird 보고 오자
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia" // Name, Origin은 Animal의 속성들임
	b.SpeedKPH = 48
	b.CanFly = false // go의 bool은 가장 앞 문자가 소문자다
	fmt.Println(b)   //{{Emu Australia} 48 false}

	// // 포함된 구조체의 속성도 그냥 .하나로 호출할 수 있음
	fmt.Println(b.Name)
	// fmt.Println("")

	// Embed라는 말의 의미는 그런 거다
	/* Bird "is" an animal 이라는 전통의 Inheritance Relationship이 아니라
	Bird "has" an animal (characteristic) 이라는 식으로 바뀐 거임
	이는 Bird = animal을 의미하지 않음 : 서로 바뀌어서(Interchangably) 사용될 수 없기 때문.
	Interface를 통해 상호 교환 가능하게 할 수 있지만 나중에 얘기함 */

	// LiteralSyntax : 선언할 때는 Embedding을 신경써야 함
	// b := Bird{
	// 	Animal:   Animal{Name: "Emu", Origin: "Australia"},
	// 	SpeedKPH: 48,
	// 	CanFly:   false,
	// }
	// fmt.Println(b.Name)

	// 공통된 속성을 만들고 싶을 때는 Interface를 이용하는 게 더 좋다
	// Embedding이 좋을 때는? 라이브러리나 웹 프레임워크 등의 base controller를 만들 때

	// Tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

// 어 근데 go 큰따옴표 작은따옴표 구분하나본데 / '를 넣으니까 Illegal Rune Syntax라고 뜸
