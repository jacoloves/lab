fn gen_message() -> &str {
    let msg = String::from("過ちは見過ごす人は美しい");
    return &msg;
}

fn main() {
    let m = gen_message();
    println!("{}", m);
}
