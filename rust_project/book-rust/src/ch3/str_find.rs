fn main() {
    let s = "隣の客はよく柿食う客だ";

    match s.find('柿') {
        Some(i) => println!("柿={}B", i),
        None => println!("柿はなし"),
    };

    match s.find("バナナ") {
        Some(i) => println!("バナナ={}B", i),
        None => println!("バナナはなし"),
    };
}
