fn smpLet() {
    let condition = true;
    let number = if condition {
        5
    } else {
        6
    };

    println!("The value of number is: {}", number);
}

fn smpElseIf() {
    let number = 6;

    if number % 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3, or 2");
    }
}

fn smpIf() {
    let number = 7;

    if number < 5 {
        println!("condition was true");
    } else {
        println!("condition was false");
    }
}

fn plus_one(x: i32) -> i32 {
    x + 2
}

fn five() -> i32 {
   5 
}

fn smpScorp() {
    let x = 5;

    let y = {
        let x = 3;
        x + 1
    };

    println!("The value of y is: {}", y);
}

fn another_function(x: i32, y: i32) {
    println!("The value of x is: {}", x);
    println!("The value of y is: {}", y);
}

fn smpArray() {
    let a = [1, 2, 3, 4, 5];

    let month = ["January", "February", "March", "April", "May", "June", "July",
                "August", "September", "October", "November", "December"];

    let first = a[0];
    let second = a[1];
}

fn smpTup2() {
    let x: (i32, f64, u8) = (500, 6.4, 1);
    
    let five_hundred = x.0;

    let six_point_four = x.1;

    let one = x.2;

    println!("The value of x is: {}", five_hundred);
    println!("The value of y is: {}, {}", six_point_four, one);
    println!("The value of z is: {}", one);
}

fn smpTup() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);

    let (x, y, z) = tup;

    println!("The value of y is: {}", y);
}

fn smpString() {
    let c = 'z';
    let z = 'ℤ';
    let heart_eyed_cat = '😻';
}

fn smpBool() {
    let t = true;

    let f: bool = false;
}

fn smp_math() {
    let x = 2.0;

    let y: f32 = 3.0;

    let sum = 5 + 10;

    let difference = 95.5 - 4.3;

    let product = 4 * 30;

    let quotient = 56.7 / 32.2;

    let remainder = 43 % 5;
}

fn shadowing() {
    let x = 5;

    let x = x + 1;
  
    let x = x * 2;

    println!("The value of x is: {}", x);
}

fn main() {
    /*
    let mut x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    println!("The value of x is: {}", x);
    */
    shadowing();
    smpTup();
    smpTup2();
    another_function(5, 6);
    smpScorp();

    let x = five();

    println!("The value of x is: {}", x);

    let y = plus_one(5);

    println!("The value of x is: {}", y);
    smpIf();
    smpElseIf();
    smpLet();
}
