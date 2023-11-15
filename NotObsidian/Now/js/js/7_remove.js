const buttons = document.querySelectorAll("p > span"); // 모든 삭제 버튼

for(let button of buttons){
    button.addEventListener("click", function () {
        this.parentNode.remove(this);
    })
}