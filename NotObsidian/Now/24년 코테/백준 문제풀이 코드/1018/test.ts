const input: string[] = require('fs').readFileSync('test.txt').toString().split('\n');

const [M, N] = input[0].split(' ').map(Number)
const matrix: string[] = input.slice(1) // 2번째 원소 ~ 끝까지 슬라이스

const Wmatrix = (row: number, col: number): number => {
    let needToFix: number = 0;
    for (let r=0;r<8;r++) {
        for (let c=0;c<8;c++) {
            const cellValue = matrix[row + r][col + c]
            if ((r + c) % 2 === 0 && cellValue === 'B') {
                needToFix += 1
            } else if ((r + c) % 2 === 1 && cellValue === 'W') {
                needToFix += 1
            }
        }
    }
    return needToFix
}

const Bmatrix = (row: number, col: number): number => {
    let needToFix: number = 0;
    for (let r=0;r<8;r++) {
        for (let c=0;c<8;c++) {
            const cellValue = matrix[row + r][col + c]
            if ((r + c) % 2 === 0 && cellValue === 'W') {
                needToFix += 1
            } else if ((r + c) % 2 === 1 && cellValue === 'B') {
                needToFix += 1
            }
        }
    }
    return needToFix
}

let ans = 1251;
for (let i=0;i<M-7;i++) {
    for (let j=0;j<N-7;j++) {
        ans = Math.min(ans, Wmatrix(i,j), Bmatrix(i, j))
    }
}

console.log(ans)