const title = document.querySelector('#title');
const author = document.querySelector('#author');
const save = document.querySelector('#save');
const bookList = document.querySelector('#booklist');

save.addEventListener("click", (e) => {
    e.preventDefault();

    // 항목 만들기
    const item = document.createElement('li');

    item.innerHTML = `
        ${title.value} - ${author.value}
        <span class="delButton">삭제</span>
    `;
    bookList.appendChild(item);

    title.value = "";
    author.value = "";

    const delButtons = document.querySelectorAll('.delButton');

    for (let delButton of delButtons) {
        delButton.addEventListener("click", function () {
            this.parentNode.parentNode.removeChild(this.parentNode);
        })
    }
});