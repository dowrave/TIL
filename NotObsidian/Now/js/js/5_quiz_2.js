
let bttn = document.querySelector('button')

function getGCD(a, b) {
    let c
    while (b !== 0) {
        c = a % b
        a = b
        b = c
    }

    return a
}

bttn.onclick = function() {
    let a = document.querySelectorAll('input')[0].value
    let b = document.querySelectorAll('input')[1].value

    ans = getGCD(a, b)

    document.querySelector('.result').innerText = ans;
}