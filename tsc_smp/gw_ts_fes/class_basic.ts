class Person {
    name: string;
    age: number;

    constructor(name: string, age: number) {
        this.name = name;
        this.age = age;
    }

    show(): string {
        return `${this.name}は${this.age}歳です。`;
    }
}

let p = new Person('rio', 30);
console.log(p.show());
