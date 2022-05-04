interface Car {
    type: string;
    run(): void;
}

let c: Car = {
    type: 'track',
    run(){
        console.log(`${this.type}が走ります`);
    }
};

c.run();
