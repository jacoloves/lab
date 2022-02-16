class Person {
    name: string;sex: string;age: number;
    constructor(name, sex) {
        this.name = name;
        this.sex = sex;
    }

    show() {
        return `${this.name}は${this.sex}です。`;
    }
}

let p = new Person('Rio', 'woman');
console.log(p.show());
