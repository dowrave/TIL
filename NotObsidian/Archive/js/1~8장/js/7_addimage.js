const today = new Date(); // 현재 날짜와 시간 정보를 담음
const hrs = today.getHours(); // 시Hour 정보를 가져옴
const container = document.querySelector("#container");

let newImg = document.createElement("img");
newImg.src = (hrs < 12) ? "images/morning.jpg" : "images/afternoon.jpg";
container.appendChild(newImg);