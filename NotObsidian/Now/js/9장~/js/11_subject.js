const member1 = ['HTML', "CSS"];
const member2 = ['CSS', '자바스크립트', '리액트'];
const member3 = ['자바스크립트', '타임스크립트'];

const subjects = [...member1, ...member2, ...member3];

const resultList = new Set();

subjects.forEach(subject => {
    resultList.add(subject);
});

// #result 영역에 resultList 표시
const result = document.querySelector('#result');

result.innerHTML = `
    <ul>
    ${[...resultList]
    .map(subject => `<li>${subject}</li>`)
    .join("")}
    </ul>
`;
