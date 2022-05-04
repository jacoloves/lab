function log(target: any, key: string, desc: PropertyDecorator) {
    let origin = desc.value;
    
    desc.value = function() {
        console.log(`${key} start...`);
        let start = Date.now();
        let result = origin.apply(this, arguments);
        let end = Date.now();
        console.log(`${key} end...`);
        console.log(`Process Time ${end - start}ms`);
        return result;
    }
}

class MyClass {
    @log
    add(x: number, y: number): number {
        let s = Date.now();
        while (Date.now()-s < 4500);

        return x+y;
    }
}

let x = new MyClass();
console.log(c.add(10,20));
