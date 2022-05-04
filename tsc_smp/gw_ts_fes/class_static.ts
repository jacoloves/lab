class Figure {
    public static PI: number = 3.14159;
    public static circle(radius: number): number {
        return radius * radius * this.PI;
    }
}

console.log(Figure.PI);
console.log(Figure.circle(5));
