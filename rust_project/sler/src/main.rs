fn clear() {
   println!("\x1b[2J");
   println!("\x1b[H");
}

fn main() {
    self::clear();
    println!("Hello, world!");
}
