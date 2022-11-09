use std::io::{BufRead, BufReader};
#[allow(unused_imports)]
use std::fs::File;

fn main() {
    let dicfile = "ejdict-hand-utf8.txt";

    let args: Vec<String> = std::env::args().collect();

    if args.len() < 2 {
        println!("[USAGE] jisyo word");
        return;
    }

    let word = &args[1];

    let fp = File::open(dicfile).unwrap();

    let reader = BufReader::new(fp);
    for line in reader.lines() {
        let line = line.unwrap();

        if line.find(word) == None { continue; }
        println!("{}", line);
    }
}
