// 사용자가 입력한 숫자를 짝수인지 홀수인지 체크하는 스크립트를 만드시오

let userNumber = parseInt(prompt("숫자를 입력하세요"));

if (userNumber !== null) {

    userNumber = parseInt(userNumber);
    (userNumber % 2 === 0) ? alert(`${userNumber} : 짝수`) : alert(`${userNumber} : 홀수`);
};