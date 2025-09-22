function countWordsInArr(arrStr, arrNum) {
  let count = 0;
  for (let word of arrStr) {
    if (arrNum.includes(word.length)) {
      count++;
    }
  }
  return count;
}

let strings = ["I", "am", "not", "a", "robot"];
let numbers = [1, 2, 3];

console.log(countWordsInArr(strings, numbers));
