const canvas = document.querySelector("canvas");
const ctx = canvas.getContext("2d");


// 얼굴 형태
ctx.scale(1, 0.7);

ctx.fillStyle = "green";
ctx.strokeStyle = "black";
ctx.beginPath();
ctx.arc(150, 150, 80, 0, Math.PI * 2, false);
ctx.fill();

// 눈
ctx.beginPath();
ctx.arc(120, 80, 20, 0, Math.PI * 2, false);

// 오른쪽 눈
ctx.moveTo(200, 80);
ctx.arc(180, 80, 20, 0, Math.PI * 2, false);

ctx.fillStyle = "white";
ctx.strokeStyle = "green";
ctx.fill();
ctx.stroke();

// 눈동자
ctx.beginPath();
ctx.arc(120, 80, 5, 0, Math.PI * 2, false);

// 오른쪽 눈동자
ctx.moveTo(200, 80);
ctx.arc(180, 80, 5, 0, Math.PI * 2, false);

ctx.fillStyle = "black";
ctx.fill();

// 입
ctx.beginPath();
ctx.arc(150, 150, 50, 0, (Math.PI / 180)*180, false);
ctx.strokeStyle = "black";
ctx.lineWidth = 3;
ctx.stroke();