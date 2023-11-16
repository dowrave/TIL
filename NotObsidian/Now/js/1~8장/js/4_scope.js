function addSum(n) {
    let sum = 0;
    for (let i = 1; i <= n; i ++){
        sum += i;
    }
    return sum;
}

const num = 3;
console.log(`${addSum(num)}`);