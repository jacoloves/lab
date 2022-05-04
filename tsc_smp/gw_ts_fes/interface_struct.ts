interface Figure {
    getArea(): number;
}

class Triangle {
    constructor(private width: number, protected height: number) {}
    getArea(): number {
        return this.width * this.height / 2;
    }
}

let t: Figure = new Triangle(10,5);
console.log(t.getArea());
