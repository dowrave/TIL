const orderButton = document.querySelector("#order");
const orderInfo = document.querySelector("#orderInfo");
const title = document.querySelector("#container > l2");

orderButton.addEventListener("click", () => {
    let newP = document.createElement("p");
    let textNode = document.createTextNode(title.innerText);
    newP.aapendChild(textNode);
    orderInfo.appendChild(newP);
});