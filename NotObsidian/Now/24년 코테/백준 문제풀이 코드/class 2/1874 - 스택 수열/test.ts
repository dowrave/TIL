const input: string[] = require('fs').readFileSync('test.txt').toString().split('\n')
const numInput: number[] = input.map((value) => Number(value.trim()))

const N = numInput[0]
const targets = numInput.slice(1)

const stack: number[] = []
const ans: string[] = []

let nowNumber = 1

for (let target of targets) {
    if (stack.length === 0) {
        stack.push(nowNumber)
        nowNumber++;
        ans.push('+')
    }
    
    while (nowNumber <= target) {
        stack.push(nowNumber)
        nowNumber++;
        ans.push('+')
    }

    if (stack[stack.length - 1] === target) {
        stack.pop()
        ans.push('-');
    }
}

if (stack.length === 0) {
    for (let i of ans) {
        console.log(i)
    }
} else {
    console.log("NO")
}