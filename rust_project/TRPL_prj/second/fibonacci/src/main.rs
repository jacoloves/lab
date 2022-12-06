use std::io;

fn main() {
    let mut f0 = 0;
    let mut f1 = 1;

    let mut f2 = 0;

    let mut cnt = 1;

    let mut fibonacci_num = String::new();

    println!("input fibonacci number");
    io::stdin()
        .read_line(&mut fibonacci_num)
        .expect("Failed to read line");

    let fibonacci_num: u32 = fibonacci_num.trim().parse().ok().unwrap();

    println!("{}", f0);

    loop {
        if cnt == fibonacci_num {
            break;
        }
        cnt += 1;
        println!("{}", f1);

        f2 = f0 + f1;

        f0 = f1;
        f1 = f2;
    }

    println!("Finish!!");
}
