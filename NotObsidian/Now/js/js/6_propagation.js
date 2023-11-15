// 모든 요소를 가져와 elements에 저장
const elements = document.querySelectorAll('*');

// 모든 요소를 순회, click 이벤트 발생 시 event.target인 태그 이름과 event.currentTarget 태그 이름을 출력한다.
for (let element of elements) {
    element.addEventListener("click", e =>
    console.log(`event.target : ${e.target.tagName}, event.currentTarget : ${e.currentTarget.tagName}`), true);
}