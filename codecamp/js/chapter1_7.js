myVar = "test";

function myLocalScope() {
  let myVar = "test2";
  console.log('inside myLocalScope', myVar);
}
myLocalScope();
console.log('outside myLocalScope', myVar);

const outerWear = "T-Shirt";
function myOutfit() {
  const outerWear = "sweater";
  return outerWear;
}
myOutfit();

let sum = 0;

function addThree() {
  sum = sum + 3;
}

function addFive() {
  sum = sum + 5;
}
addThree();
addFive();
console.log(sum);

let processed = 0;
function processArg(num) {
  return (num + 3) / 5;
}
processed = processArg(7);

function nextInLine(arr, item) {
  arr.push(item);
  var removed = arr.shift();

  return removed;
}
const testArr = [1, 2, 3, 4, 5];
console.log("Before: " + JSON.stringify(testArr));
console.log(nextInLine(testArr, 6));
console.log("After: " + JSON.stringify(testArr));
