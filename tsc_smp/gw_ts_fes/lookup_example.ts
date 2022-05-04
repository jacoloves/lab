interface Product {
    name: string;
    price: number;
}

function getProp<T, K extends keyof T>(obj: T, name: K) {
    return obj[name];
}

let p = { name: 'obento', price: 500 };
console.log(getProp(p, 'name'));
console.log(getProp(p, 'price'));
