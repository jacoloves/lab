use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        println!("入力ファイルを指定してください。");
        return;
    }

    let filename = &args[1];
    let text = fs::read_to_string(filename).unwrap();
    println!("{}", text);
}
