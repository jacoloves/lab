use std::fs;

fn main() {
    let files = fs::read_dir(".").expect("illegal path");
    for ent in files {
        let entry = ent.unwrap();
        let path = entry.path();
        let fname = path.to_str().unwrap_or("illegal file name");
        println!("{}", fname);
    }
}
