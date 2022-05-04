class MyClass {
    #data = 10;
    show(): void {
        console.log(`値は${this.#data}です。`);
    }
}

let c = new MyClass();
c.show();
