canvas = document.querySelector('canvas');
ctx = canvas.getContext('2d');

ctx.fillstyle = '#ccc';
ctx.fillRect(10, 10, 100, 100);

ctx.save(); // 일단 저장 먼저

// translate : 원점 이동
// ctx.translate(50, 50); // 원점 이동
// ctx.fillStyle = '#222';
// ctx.fillRect(10, 10, 100, 100);
// ctx.fillStyle = 'red';
// ctx.fillRect(50, 50, 80, 20);
// ctx.restore();

// rotate : 마름모 사각형
// ctx.translate(10, 10);
// ctx.rotate(45 * Math.PI / 180);
// ctx.strokeRect(0, 0, 100, 100);
// ctx.translate(-10, -10);

// scale : 도형 확대/축소
ctx.scale(3, 2);
ctx.strokeRect(20, 70, 100, 50); // strokeRect(60, 140, 300, 100);

ctx.restore();

ctx.strokeRect(200, 50, 100, 50);