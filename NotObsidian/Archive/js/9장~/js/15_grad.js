canvas1 = document.querySelector('#linear')
canvas2 = document.querySelector('#radial')
ctx1 = canvas1.getContext('2d')
ctx2 = canvas2.getContext('2d')

let linGrad = ctx1.createLinearGradient(0, 0, 0, 200);
linGrad.addColorStop(0, "#000"); // 시작 위치 : 검은색
linGrad.addColorStop(0.6, "#fff"); // 0.6 위치 : 흰색
linGrad.addColorStop(1, "#eee"); // 끝나는 위치 : 회색

ctx1.fillStyle = linGrad;
ctx1.fillRect(0, 0, 100, 200);

ctx2.shadowColor = "#ccc";
ctx2.shadowOffsetX = 15;
ctx2.shadowOffsetY = 10;
ctx2.shadowBlur = 10;

let radGrad = ctx2.createRadialGradient(55, 60, 10, 80, 90, 100);
radGrad.addColorStop(0, "white");
radGrad.addColorStop(0.4, "yellow");
radGrad.addColorStop(1, "orange");

ctx2.fillStyle = radGrad;
ctx2.arc(100, 100, 80, 0, Math.PI * 2, false);
ctx2.fill();