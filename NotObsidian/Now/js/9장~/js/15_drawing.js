const canvas = document.querySelector("#canvas");
const toolbar = document.querySelector('#toolbar');

canvas.width = window.innerWidth;
canvas.height = window.innerHeight - toolbar.offsetHeight;

// 마우스 커서 위치 계산 : 툴바까지 계산해야 하므로 캔버스의 왼쪽과 윗쪽으로부터 커서의 위치를 구한다.
const canvasOffsetX = canvas.offsetLeft;
const canvasOffsetY = canvas.offsetTop; 

// 그리는 중 & 시작 위치 변수 & 선 굵기
const ctx = canvas.getContext('2d');

let isDrawing = false;
let startX, startY;
let lineWidth = 2;

// 툴바
toolbar.addEventListener('change', e => {
    if (e.target.id === 'stroke') {
        ctx.strokeStyle = e.target.value;
    }
    if (e.target.id === 'lWidth') {
        lineWidth = e.target.value;
    }
});

// 지우기 버튼
toolbar.addEventListener("click", e => {
    if (e.target.id === "reset") {
        ctx.clearRect(0, 0, canvas.width, canvas.height);
    }
});

// 그림 그리기
canvas.addEventListener("mousedown", e => {
    isDrawing = true;
    startX = e.clientX;
    startY = e.clientY;
})

// 선 그리기 : 툴바 영역 제외
canvas.addEventListener("mousemove", e => {
    if (!isDrawing) return;
    ctx.lineWidth = lineWidth;
    ctx.lineCap = 'round';
    ctx.lineTo(e.clientX, e.clientY - canvasOffsetY);
    ctx.stroke();
})

// 손 뗴면 끝내기
canvas.addEventListener("mouseup", () => {
    isDrawing = false;
    ctx.beginPath();
})

