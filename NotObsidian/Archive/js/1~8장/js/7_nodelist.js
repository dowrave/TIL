let tsNode = document.createElement('p')
let tsTextNode = document.createTextNode('TypeScript')
let basisNode = document.querySelectorAll('p')[0]
const items = document.querySelectorAll('li'); 

tsNode.appendChild(tsTextNode);
document.body.insertBefore(tsNode, basisNode);

// n번째 li 요소 삭제하기

for(item of items) {
    item.addEventListener("click", function () { // 항목 클릭시 함수 실행
        this.parentNode.removeChild(this); 
    });
}
