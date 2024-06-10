canvas = document.querySelector('canvas');
ctx = canvas.getContext('2d');

let img = new Image();
img.onload = function() {
	ctx.drawImage(img, 0, 0, canvas.width, canvas.height)
};
img.src = 'cat.jpg';

// 클리핑
ctx.beginPath();
ctx.arc(300, 200, 150, 0, Math.PI * 2, false);
ctx.clip();