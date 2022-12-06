use std::io;

fn main() {
    println!("Please input celsius");

    let mut celsius_num = String::new();

    io::stdin()
        .read_line(&mut celsius_num)
        .expect("Failed to read line");

    let celsius_num: u32 = match celsius_num.trim().parse() {
        Ok(num) => num,
        Err(_) => 0,
    };

    println!("Please input fahrenheit");
    let fahrenheit_num: u32 = (9 * celsius_num) / 5 + 32;

    println!("1: celsius: {}", celsius_num);
    println!("1: fahrenheit: {}", fahrenheit_num);

    let mut fahrenheit_num_second = String::new();

    io::stdin()
        .read_line(&mut fahrenheit_num_second)
        .expect("Failed to read line");

    let fahrenheit_num_second: u32 = match fahrenheit_num_second.trim().parse() {
        Ok(num) => num,
        Err(_) => 0,
    };

    let celsius_num_second: u32 = 5 * (fahrenheit_num_second - 32) / 9;

    println!("2: fahrenheit: {}", fahrenheit_num_second);
    println!("2: celsius: {}", celsius_num_second);
}
