

function judge(){
    let isValid = false
    let inputNumber
    
    while (!isValid){
        inputNumber = parseInt(prompt("정수를 입력하세요"));

        if(inputNumber === 0) {
            alert('0입니다.')
            isValid = true
        }
        else if (inputNumber % 2 === 0){
            alert("짝수입니다.")
            isValid = true
        }
        else if (inputNumber % 2 === 1){
            alert("홀수입니다.")
            isValid = true
        }
        else{
            alert("다시 입력해주세요.")
        }
    }
}

judge()