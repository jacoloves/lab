interface Person {
    name: string;
    age: number;
}

function greet(p: Person): void {
    console.log(`Hello, ${p.name}!`);
}

let pl = {
    name: 'sakura',
    age: 2,
    gender: 'female',
};

greet(pl);
