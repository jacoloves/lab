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
    protected clazz: string;
    constructor(name: string, age: number, clazz: string) {
        super(name,age);
        this.clazz = clazz;
    }

    show(): string {
        return super.show() + `${this.clazz}です。`;
    }
}

let p = new BusinessPerson('rio', 30, 'manager');
console.log(p.show());
