fn transFah(num: f64) -> f64 {
    (num * 1.8) + 32.0
}

fn transCel(num: f64) -> f64 {
    (num - 32.0) / 1.8
}

fn main() {
    let x = 32.0;

    println!("Fahrenheit is {}", transFah(x));
    
    let y = 57.6;

    println!("Celsius is {}", transCel(y));

}
