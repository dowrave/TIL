const canvas = document.querySelector("canvas");
const ctx = canvas.getContext('2d');

canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

function Circle(x, y, radius, color){
    this.x = x;
    this.y = y;
    this.radius = radius;
    this.color = color;
	this.dx = Math.floor(Math.random()  * 4) + 1
	this.dy = Math.floor(Math.random()  * 4) + 1

    this.draw = function() {
        ctx.beginPath();
        ctx.fillStyle = this.color;
        ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2, false);
        ctx.fill();
    }

	this.animate = function() {
	    if (this.x + this.radius > canvas.width || this.x - this.radius < 0 ){
	        this.dx = - this.dx
	    }

	    if (this.y + this.radius > canvas.height || this.y - this.radius < 0 ){
	        this.dy = - this.dy
	    }    
	
		this.x += this.dx;
	    this.y += this.dy;
	    
	    this.draw();
	}
}

// const circleOne = new Circle(100, 100, 50, "red");
// const circleTwo = new Circle(200, 200, 20, "blue");

// circleOne.draw();
// circleTwo.draw();

// 무작위로 20개의 원 그리기
// 랜덤으로 넣을 값은 x, y, radius, color - r, g, b로 총 6가지임
const objs = []

for (let i = 0;i < 20; i++){
    let radius = Math.random() * 100;
    let random_x = Math.floor(Math.random() * canvas.width)
    let random_y = Math.floor(Math.random() * canvas.height)
    let r = Math.floor(Math.random() * 256);
    let g = Math.floor(Math.random() * 256);
    let b = Math.floor(Math.random() * 256);
    let alpha = 0.5;

    objs.push(new Circle(random_x, random_y, radius, `rgba(${r}, ${g}, ${b}, ${alpha})`))
}

function update() {

	// 1.캔버스 지우기
	ctx.clearRect(0, 0, canvas.width, canvas.height)
	
	// 2. 반복문 - objs의 각 요소마다 애니메이션 실행하기
	for (let i=0; i<objs.length; i++){
		let obj = objs[i];
		obj.animate();
	}
	requestAnimationFrame(update)
};

update();

