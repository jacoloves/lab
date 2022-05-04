class MyClass {
    constructor(private _value: number) {}
    get value(): number {
        return this._value;
    }

    plus(value: number): this {
        this._value += value;
        return this;
    }

    minus(value: number): this {
        this._value -= value;
        return this;
    }
}

let clazz = new MyClass(10);
console.log(clazz.plus(10).minus(5).value);
