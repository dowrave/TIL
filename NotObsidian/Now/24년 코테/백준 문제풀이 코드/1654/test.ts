const input: string[] = require('fs').readFileSync('dev/stdin').toString().trim().split('\n');
const [K, N] = input[0].split(' ').map(Number);
const arr: number[] = input.slice(1).map(Number);

let start = 1
let end = Math.max(...arr)
let ans = 0

while (start <= end) {
    const mid = Math.floor((start + end) / 2);
    let count = 0;

    for (const num of arr) {
        count += Math.floor(num / mid);
    }

    // N보다 작다 = 길이가 너무 길다
    if (count < N) {
        end = mid - 1
    } 
    // N보다 길다 또한 N개를 만들었다에 포함되므로
    else {
        ans = mid
        start = mid + 1
    }
}

console.log(ans)

