abstract class Figure {
    constructor(protected width: number, protected height: number) {
        console.log('Triangle value set!!');
    }
    abstract getArea(): number; 
}

class Triange extends Figure {
    getArea(): number {
        return this.width * this.height / 2;
    }
}

let t = new Triange(10, 5);
console.log(t.getArea());
