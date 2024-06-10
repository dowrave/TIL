// 논리는 맞는 것 같은데 틀렷습니다가 뜸. 일단 넘김.
const N: string = require('fs').readFileSync('test.txt').toString().trim()
const number = parseInt(N)

const iter_factorial = ((x: number) => {
    let ans = 1
    for (let i=1; i < x + 1; i++) {
        ans *= i
    }
    return ans 
})

let ans = iter_factorial(number);
let count = 0

while (true) {
    if (ans % 10 === 0) {
        count++;
        ans = Math.floor(ans / 10)
    } else {
        break
    }
}

console.log(count)