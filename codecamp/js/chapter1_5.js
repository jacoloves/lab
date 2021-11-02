const myArray = ["test", 3];

const myArray = [["test", 4]];

const myArray = [50, 60, 70];
const myData = myArray[0];

const myArray = [18, 64, 99];
myArray[0] = 45;

const myArray = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
  [[10, 11, 12], 13, 14],
];
const myData = myArray[2][1];

const myArray = [["John", 23], ["cat", 2]];
myArray.push(["dog", 3]);

const myArray = [["John", 23], ["cat", 2]];
const removedFromMyArray = myArray.pop();

const myArray = [["John", 23], ["dog", 3]];
const removedFromMyArray = myArray.shift();

const myArray = [["John", 23], ["dog", 3]];
myArray.shift();
myArray.unshift(["Paul", 35]);
