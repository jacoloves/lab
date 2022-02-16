// $(function(){
//     $('.hoge').css('backgrounf-color', 'Red');
// });

function triangle(base: number, height: number): number {
    return base * height / 2;
}
console.log(triangle(10, 5));

// let triangle2(base: number, height: number): number {
//     return base * height / 2;
// }
// console.log(triangle(10, 5));

let triangle3: (base: number, height: number) => number =
    function (base: number, height: number): number {
        return base * height / 2;
    }

let triangle4 = (base: number, height: number): number => {
    return base * height / 2;
};
console.log(triangle4(10, 5));

let Counter = function() {
    let _this = this;
    _this.count = 0;
    setInterval(function() {
        _this.count++;
    }, 1000);
};

function showTime(time: Date = new Date()): string {
    if (time === undefined) {
        return '現在時刻: ' + (new Date()).toLocaleString();
    } else {
        return '現在時刻: ' + time.toLocaleString();
    }
}

console.log(showTime());

function sum(...values: number[]) {
    let result = 0;
    for (let value of values) {
        result += value;
    }
    return result;
}
console.log(sum(1, 5, -8, 10));

function show(value: number): void;
function show(value: boolean): void;
function show(value: any): void {
    if (typeof value === 'number') {
        console.log(value.toFixed(0));
    } else {
        console.log(value ? 'true' : 'false');
    }
}

show(10.358);
show(false);

function square(value: number): number | boolean {
    if (value < 0) {
        return false; 
    } else {
        return Math.sqrt(value);
    }
}
console.log(square(9));
console.log(square(-9));