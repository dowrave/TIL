const bttn = document.querySelector("button");
const popWidth = 600;
const popHeight = 500;


function openPopup() {
    let left = (screen.availWidth - popWidth) / 2;
    let top = (screen.availHeight - popHeight) / 2;
    window.open("../8_poppedup.html", "이벤트 팝업", `width=${popWidth} height=${popHeight} top=${top} left=${left}`)
}

// bttn.onclick = openPopup;
bttn.addEventListener("click", openPopup);