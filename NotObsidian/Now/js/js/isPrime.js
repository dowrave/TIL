let inputNumber = parseInt(prompt("자연수를 입력하세요."));
let isPrime = '소수';

// inputNumber가 소수인지 체크하려면
// 1. 1은 그냥 아니죠? false
// 2. 2부터 숫자를 1씩 올리면서 이 숫자로 나눠 떨어지는가를 체크한다.
// 3. inputNumber의 제곱근까지 2.를 반복한다.
if (inputNumber > 1) {
    for (i=2; i < inputNumber ** (1/2); i++) {
        if (inputNumber % i === 0) {
           isPrime = '안소수';
           break
        }
    }
}
else {isPrime = '안소수'};


document.write(`${inputNumber} : ${isPrime}입니다.`)