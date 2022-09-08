
fn main() {
    let s: &str = "test";
    let str_len: usize = s.chars().count();
    let i: i32 = 7;
    let space_len: i32 = i - str_len as i32;
    let mut space = "".to_string();
    for _ in 1..=space_len {
        space.push_str(" ");
    }
    println!("{}{}", space, s);
}
