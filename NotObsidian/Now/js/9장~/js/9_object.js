// 객체 선언
let book1 = {
        'title' : '제목',
        'pages' : 648
}

let book2 = {
        title : "제목입니다",
        pages : 335
}

div_tag = document.querySelector('div');
div_tag.innerText = book1.title;