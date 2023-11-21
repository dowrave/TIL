let students = ['kim', 'lee', 'park']
console.log(`현재 students : ${students}`)

// 스토리지에 students 키로 배열 저장
localStorage.setItem("students", JSON.stringify(students));

// 새로운 값 choi 추가하기
let localData;
if(localStorage.getItem("students") === null) {
    localData = [];
} else {
    localData = JSON.parse(localStorage.getItem("students"));
}

localData.push("choi");
localStorage.setItem('students', JSON.stringify(localData));
console.log(`추가 후 students : ${localData}`)

// 로컬 스토리지의 값 `lee` 삭제하기
const indexOfValue = localData.indexOf('lee');
localData.splice(indexOfValue, 1); // 인덱스에 해당하는 값부터 1개를 삭제함
localStorage.setItem('students', JSON.stringify(localData));
console.log(`삭제 후 students : ${localData}`)

localStorage.removeItem('students');