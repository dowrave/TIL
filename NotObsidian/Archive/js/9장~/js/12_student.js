let xhr = new XMLHttpRequest();

xhr.open('GET', 'student.json');
xhr.send();

xhr.onreadystatechange = function() {
    if ( xhr.readystate === 4 && xhr.status == 200 ) {
        let students = JSON.parse(xhr.responseText);
        document.querySelector("#result").innerText = students;
        renderHTML(students);
    }
}

function renderHTML(contents) {
    let output = "";
    for (let content of contents) {
        output += `
            <h2>${content.name}</h2>
            <ul>
                <li>전공 : ${content.major}</li>
                <li>학년 : ${content.grade}</li>
            </ul>
            </hr>
        `;
    }
    document.querySelector('#result').innerHTML = output;
}