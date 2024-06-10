const todoInput = document.querySelector('#todo-input');
const addButton = document.querySelector('#add-button');
const todoList = document.querySelector('#todo-list');

document.addEventListener('DOMContentLoaded', getLocal); 
addButton.addEventListener('click', addTodo);
todoList.addEventListener('click', manageTodo);

// 이미 로컬 스토리지에 데이터가 있다면 불러온다.
function getLocal() {
    let todos;

    if (localStorage.getItem('todos') === null){
        todos = [];
    } else {
        todos = JSON.parse(localStorage.getItem('todos'));
    }

    todos.forEach(function(todo) {

        const newDiv = document.createElement('div');
        newDiv.classList.add('todo');
        const newTodo = document.createElement("li");
        newTodo.innerText = todo;
        newTodo.classList.add('todo-content');
        newDiv.appendChild(newTodo); 

        const completeButton = document.createElement('button');
        completeButton.innerText = '완료';
        completeButton.classList.add('complete-button');
        newDiv.appendChild(completeButton);

        const deleteButton = document.createElement('button');
        deleteButton.innerText = '삭제';
        deleteButton.classList.add('delete-button');
        newDiv.appendChild(deleteButton);

        todoList.appendChild(newDiv); 
    })
}


// 추가된 데이터를 형식에 따라 나열함
function addTodo(e) {
    e.preventDefault();
    saveToLocal(todoInput.value)

    // 코드 쓰면서 느끼는 거 : 새로운 요소를 추가했는데, 어디에 추가됨? 
    const newDiv = document.createElement('div');
    newDiv.classList.add('todo');
    const newTodo = document.createElement("li");
    newTodo.innerText = todoInput.value;
    newTodo.classList.add('todo-content');
    newDiv.appendChild(newTodo); // 일단 새로 추가한 div 밑에 li가 들어감

    // 내용 오른쪽에 버튼 추가
    const completeButton = document.createElement('button');
    completeButton.innerText = '완료';
    completeButton.classList.add('complete-button');
    newDiv.appendChild(completeButton);

    const deleteButton = document.createElement('button');
    deleteButton.innerText = '삭제';
    deleteButton.classList.add('delete-button');
    newDiv.appendChild(deleteButton);

    todoList.appendChild(newDiv); // 
    todoInput.value = ""; // 입력창 초기화

    
}

// 로컬 스토리지에 저장
function saveToLocal(todo) {
    let todos;

    // 새로 입력된 내용이 있다면 불러와서 객체화 -> 추가 -> 저장
    if (localStorage.getItem('todos') === null){
        todos = [];
    } else {
        todos = JSON.parse(localStorage.getItem('todos'));
    }

    todos.push(todo);
    localStorage.setItem('todos', JSON.stringify(todos));
}

// 완료, 삭제 버튼 클릭 시 동작 정의
function manageTodo(e) {
    const whichButton = e.target.classList[0]; // 클릭한 부분의 class 명을 가져옴
    if (whichButton === 'complete-button'){
        const todo = e.target.parentElement;
        todo.children[0].classList.toggle('completed');
    } else if (whichButton == 'delete-button') {
        const todo = e.target.parentElement;
        removeLocal(todo); // 로컬 스토리지에서 todo에 해당하는 값을 제거함
        todo.remove(); // 부모 태그 `div` 를 제거함
    }
}

// 삭제버튼 동작 정의
function removeLocal(todo) {
    let todos;
    if (localStorage.getItem('todos') === null) {
        todos = [];
    } else {
        todos = JSON.parse(localStorage.getItem('todos'));
    }
    // console.log(todo);
    const index = todos.indexOf(todo.children[0].innerText); // [0]인 이유는 div의 가장 위 태그인 <li> 내의 innerText에 값이 있기 때문
    // console.log(index);
    todos.splice(index, 1); // index번째 요소를 삭제한다
    localStorage.setItem('todos', JSON.stringify(todos));
}