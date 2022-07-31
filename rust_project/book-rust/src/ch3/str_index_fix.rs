fn main() {
    let s = "Hello";
    let ch = s.chars().nth(0).unwrap();
    println!("{}", ch);

    let ch = s.chars().nth(2).unwrap();
    println!("{}", ch);
}
