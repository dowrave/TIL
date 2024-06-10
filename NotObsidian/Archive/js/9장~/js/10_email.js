const email = document.querySelector("#useremail");
const button = document.querySelector("button");
const result = document.querySelector("#result");

button.addEventListener("click", function() {

    if (email.value !== "") {

        let splitAddress = email.value.split('@')
        // let concealedID = splitAddress[0].slice(0, 3);
        let half = Math.floor(splitAddress[0].length / 2)
        let concealedID = splitAddress[0].slice(0, half);

        result.innerText = `${concealedID}...@${splitAddress[1]}`
        
        email.value = ""; // 입력 창 초기화
    }
});