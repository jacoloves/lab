fn smpRev() {
    for number in (1..4).rev() {
        println!("{}!", number);
    }
    println!("LIFTOFF!!!");
}

fn smpArray2() {
    let a = [10, 20, 30, 40, 50];

    for element in a.iter() {
        println!("the value is: {}", element);
    }
}

fn smpArray() {
    let a = [10, 20, 30, 40, 50];
    let mut index = 0;

    while index < 5 {
        println!("the value is: {}", a[index]);

        index = index + 1;
    }
}

fn smpLoop() {
    let mut number = 3;

    while number != 0 {
        println!("{}!", number);

        number = number + 1;
    }

    println!("LIFTOFF!!!");
}

fn main() {
    /*
    loop {
        println!("again!");
    }
    */
//    smpLoop();
    smpArray();
    smpArray2();
    smpRev();
}
