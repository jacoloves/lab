const myList = [["Chocolate Bar", 15], ["tre", 56], ["ytrf", 43], ["tr", 454], ["trtr", 34]];

function reusableFunction() {
  console.log("Hi World");
}
reusableFunction();

function functionWithArgs(param1, param2) {
  console.log(param1 + param2);
}
functionWithArgs(1, 2);
functionWithArgs(7, 9);

function timesFive(param) {
  return 5 * param;
}
const answer1 = timesFive(5);
const answer2 = timesFive(2);
const answer3 = timesFive(0);

const myGlobal = 10;


function fun1() {
  // Assign 5 to oopsGlobal Here
  oopsGlobal = 5;
}

function fun2() {
  var output = "";
  if (typeof myGlobal != "undefined") {
    output += "myGlobal: " + myGlobal;
  }
  if (typeof oopsGlobal != "undefined") {
    output += " oopsGlobal: " + oopsGlobal;
  }
  console.log(output);
}
