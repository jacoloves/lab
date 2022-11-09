fn main() {
    let g1 = String::from("過ちは見過ごす人は美しい");
    show_message(&g1);
    println!("{}", g1);
}

fn show_message(message: &String){
    println!("{}", message);
}
