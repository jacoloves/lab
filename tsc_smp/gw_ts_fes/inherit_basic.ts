class Person {
    protected name: string;
    protected age: number;

    constructor(name: string, age: number){
        this.name = name;
        this.age = age;
    }
    show(): string {
        return `${this.name}は${this.age}歳です。`;
    }
}

class BusinessPerson extends Person {
    work(): string {
        return `${this.name}はテキパキ働きます。`;
    }
}

let p = new BusinessPerson('rio', 30);
console.log(p.show());
console.log(p.work());
