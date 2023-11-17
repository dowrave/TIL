const string = prompt("영 소문자로 된 문자열을 입력하세요.");

const firstCh = string[0].toUpperCase();
const remainStr = string.slice(1);
const result = firstCh + remainStr;

document.write(result);