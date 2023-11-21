const canvas = document.querySelector("canvas");
const ctx = canvas.getContext('2d');

canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

const circle = {
	x: 100,
	y: 100,
	radius : 30,
	dx: 4, // 가로로 움직일 크기
	dy: 4, // 세로로 움직일 크기
	color: "#222"
}

function drawCircle() {
	ctx.beginPath();
	ctx.arc(circle.x, circle.y, circle.radius, Math.PI * 2, false);
	ctx.fillStyle = circle.color;
	ctx.fill();
}


function move() {

    // 원의 이동 구현(캔버스를 지운다 -> 새 점을 찍는다)
    // ctx.clearRect(0, 0, canvas.width, canvas.height); 

    // 흔적 남기기
    ctx.fillStyle = `rgba(255, 255, 255, 0.3)`;
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    // ctx.clearRect(0, 0, canvas.width, canvas.height); 
    drawCircle();

    if (circle.x + circle.radius > canvas.width || circle.x - circle.radius < 0 ){
        circle.dx = - circle.dx
    }

    // 세로 방향 이동도 구현하자
    if (circle.y + circle.radius > canvas.height || circle.y - circle.radius < 0 ){
        circle.dy = - circle.dy
    }    

	circle.x += circle.dx;
    circle.y += circle.dy;


	requestAnimationFrame(move);

}

move();