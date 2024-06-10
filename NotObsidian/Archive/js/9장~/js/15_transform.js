const canvas = document.querySelector("canvas");
const ctx = canvas.getContext("2d");
const button = document.querySelector("button");
const origin = {x:200, y:200}; // 옮길 원점 좌표
const alpha = 0.7;

ctx.fillStyle = "#ccc";
ctx.fillRect(origin.x, origin.y, 100, 100);

// 회전시키기 버튼 클릭시 30도 회전
button.onclick = function() {
    let color = randomRGB();

    ctx.globalAlpha = alpha;
    ctx.translate(origin.x, origin.y);
    ctx.rotate(Math.PI / 180 * 30);
    ctx.fillStyle = color;
    ctx.fillRect(0, 0, 100, 100);
    ctx.translate(-origin.x, -origin.y); 
}

// 색상 무작위로 바꾸기
function randomRGB() {
    let red = Math.floor(Math.random() * 256);
    let green = Math.floor(Math.random() * 256);
    let blue = Math.floor(Math.random() * 256);
    return `rgb(${red}, ${green}, ${blue})`
};