function merge<T, R>(obj1: T, obj2: R): T & R {
    let result = <T & R>{};
    for (let key in obj1) {
        (<any>result)[key] = obj1[key];
    }

    for (let key in obj2) {
        (<any>result)[key] = obj2[key];
    }

    return result;
}

let book = {
    title: 'JavaScript 本格入門',
    price: 2980,
    toString() {
        return `${this.title} ${this.price}`;
    }
};

let logger = {
    debug() {
        console.log(this.toString());
    }
};

let m = merge(book, logger);
console.log(m);
