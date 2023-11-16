let inputNumber = parseInt(prompt("자연수를 입력하세요..."));
let sum = 0;

for (i = 1;i<=inputNumber; i++) {
    if (i % 2 === 0){
        sum += i
        document.write(`${i}----${sum}<br>`);
    }
}

