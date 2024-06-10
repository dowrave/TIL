
const title = document.querySelector("#user");

// 1번쨰로 class = title인 부분을 클릭하면, clicked 스타일이 활성화된다. 
title.addEventListener("click", function() {
	this.classList.toggle("clicked");
});