fn main() {
    let pr = "知恵は武器よりも価値がある。";

    let sub3: String = pr.chars().take(2).collect();
    println!("先頭２文字: {}", sub3);
    let pr_chars: Vec<char> = pr.chars().collect();

    let sub_chars = &pr_chars[3..=4];

    let sub4: String = sub_chars.into_iter().collect();
    println!("4-5文字目: {}", sub4);
}
