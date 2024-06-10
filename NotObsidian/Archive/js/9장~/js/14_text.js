canvas = document.querySelector('canvas');
ctx = canvas.getContext('2d');

ctx.globalAlpha = 0.5;

ctx.font = "italic 60px serif";
ctx.fillText('HELLO', 50, 70);
ctx.font = "bold 60px sans-serif"
ctx.strokeText('HELLO', 50, 150);