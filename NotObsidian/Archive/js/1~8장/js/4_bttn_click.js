const bttn = document.querySelector("button"); // 버튼 요소를 가져옴

function display(){
	alert("클릭했습니다.");
}

bttn.addEventListener("click", display) // 버튼 클릭 시 display 함수 실행