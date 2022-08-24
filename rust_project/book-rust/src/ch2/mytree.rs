use std::{env, path};

fn main() {
    let args: Vec<String> = env::args().collect();

    let mut taget_dir = ".";
    if args.len() >= 2 {
        taget_dir = &args[1];
    }

    let target = path::PathBuf::from(taget_dir);
    println!("{}", taget_dir);
    tree(&target, 0);
}

fn tree(target: &path::PathBuf, level: isize) {
    let files = target.read_dir().expect("not exist path");

    for ent in files {
        let path = ent.unwrap().path();

        for _ in 1..=level {
            print!("|  ");
        }

        let fname = path.file_name().unwrap().to_string_lossy();
        if path.is_dir() {
            println!("|-- <{}>", fname);
            tree(&path, level+1);
            continue;
        }

        println!("|-- {}", fname);
    }
}
