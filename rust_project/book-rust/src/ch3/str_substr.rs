fn main() {
    let pr = "知恵は武器よりも価値がある";

    let mut sub1 = String::new();
    for (i, c) in pr.chars().enumerate() {
        if i < 2 { sub1.push(c); continue; }
        break;
    }
    println!("先頭２文字: {}", sub1);

    let mut sub2 = String::new();
    for (i, c) in pr.chars().enumerate() {
        if 3 <= i && i <= 4 { sub2.push(c); }
    }
    println!("4-5文字目: {}", sub2);
}
