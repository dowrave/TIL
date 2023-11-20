const canvasRect = document.querySelector('#rectangle');
const canvasTri = document.querySelector('#triangle');
const canvasCir = document.querySelector('#circle');
const canvasEll = document.querySelector('#ellipse');
const canvasWav = document.querySelector('#wave');
const canvasP2d = document.querySelector('#path2d');


const canvas = document.querySelectorAll('canvas');

const ctx1 = canvasRect.getContext("2d"); // 2D 콘텍스트
const ctx2 = canvasTri.getContext("2d");
const ctx3 = canvasCir.getContext("2d");
const ctx4 = canvasEll.getContext("2d");
const ctx5 = canvasWav.getContext("2d");
const ctx6 = canvasP2d.getContext("2d");


canvas.forEach((eachCanvas) => {
    eachCanvas.width = window.innerWidth / 3;
    eachCanvas.height = window.innerHeight / canvas.length;
});
// canvas.width = window.innerWidth;
// canvas.height = window.innerHeight;

// ctx1 : 사각형
ctx1.fillStyle = "rgb(200, 0, 0)";
ctx1.strokeStyle = "black";
ctx1.fillRect(10, 10, 200, 100);
ctx1.strokeRect(10, 10, 200, 100);
ctx1.fillStyle = "blue";
ctx1.fillRect(50, 50, 120, 100);
ctx1.clearRect(70, 80, 80, 45);


// ctx2 : 삼각형
ctx2.beginPath();
ctx2.moveTo(50, 50);
ctx2.lineTo(150, 100);
ctx2.lineTo(50, 150); // 이전 경로의 끝점에서 시작함
ctx2.closePath();
ctx2.stroke();

ctx2.beginPath();
ctx2.moveTo(150, 100);
ctx2.lineTo(250, 50);
ctx2.lineTo(250, 150);
ctx2.closePath();
ctx2.fillStyle = 'rgb(0, 200, 0)';
ctx2.stroke(); // 테두리 추가
ctx2.fill();

// ctx3 : 원
ctx3.fillStyle = "yellow";
ctx3.strokeStyle = "red";

ctx3.beginPath();
ctx3.arc(200, 150, 100, 0, Math.PI * 2, true); // (200, 150)에서 반지름 100인 원. 방향은 상관 없다.
ctx3.closePath();
ctx3.fill();
ctx3.stroke();

// ctx4 : 타원
ctx4.strokeStyle = "red";
ctx4.beginPath();
ctx4.ellipse(200, 70, 80, 50, 0, 0, Math.PI * 2);
ctx4.stroke();

ctx4.strokeStyle = 'blue';
ctx4.beginPath();
ctx4.ellipse(150, 200, 80, 50, (Math.PI / 180) * -30, 0, Math.PI * 2);
ctx4.stroke();
ctx4.closePath();

ctx4.strokeStyle = 'green';
ctx4.scale(1, 0.7);
ctx4.beginPath()
ctx4.arc(300, 150, 80, 0, Math.PI * 2, true);
ctx4.stroke();


ctx4.beginPath()
ctx4.arc(300, 150, 30, 0, Math.PI * 2, false);
ctx4.stroke();
ctx4.closePath();

// ctx5 : 2차 베지에 곡선 - 물결
ctx5.strokeStyle = "black";
ctx5.beginPath();
ctx5.moveTo(50, 100);
ctx5.quadraticCurveTo(100, 50, 150, 100);
ctx5.quadraticCurveTo(200, 150, 250, 100);
ctx5.quadraticCurveTo(300, 50, 350, 100);
ctx5.stroke();

// ctx5 : 3차 베지에 곡선 예제
ctx5.strokeStyle = "red";
ctx5.beginPath();
ctx5.moveTo(50, 100)
ctx5.bezierCurveTo(90, 250, 310, 10, 350, 100);
ctx5.stroke();


// ctx6 : path2D
let triangle = new Path2D();
triangle.moveTo(100, 100);
triangle.lineTo(300, 100);
triangle.lineTo(200, 260);
triangle.closePath();

let circle = new Path2D();
circle.arc(200, 155, 50, 0, Math.PI * 2);

ctx6.fillStyle = 'green';
ctx6.stroke(triangle);
ctx6.fill(circle);