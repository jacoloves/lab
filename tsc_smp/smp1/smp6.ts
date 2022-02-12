$(function(){
    $('.hoge').css('backgrounf-color', 'Red');
});

function triangle(base: number, height: number): number {
    return base * height / 2;
}
console.log(triangle(10, 5));

let triangle2(base: number, height: number): number {
    return base * height / 2;
}
console.log(triangle(10, 5));

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
    
}