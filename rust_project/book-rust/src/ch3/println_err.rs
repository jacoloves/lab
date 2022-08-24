fn main() {
    let s = "test println! error".to_string();
    echo(s);
    println!("{}", s);
}

fn echo(s: String) {
    println!("{}", s);
}
