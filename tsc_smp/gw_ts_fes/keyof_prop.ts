function getProp(obj: Product, name: keyof Product) {
    return obj[name];
}

let p = {name: 'bento', price: 500} ;
console.log(getProp(p, 'name'));
console.log(getProp(p, 'price'));
