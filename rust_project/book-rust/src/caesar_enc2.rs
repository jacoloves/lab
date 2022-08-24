fn encrypt(text: &str, shift: i16) -> String {
    let a = 'A' as i16;
    let is_az = |c| 'A' <= c && c <= 'Z';
    let conv = |c| (((c-a+shift+26)%26+a) as u8)as char;
    let encl = |c| if is_az(c) { conv(c as i16)} else { c };
    return text.chars().map(|c| encl(c)).collect();
}

fn main() {
    let enc = encrypt("I LOVE YOU", 3);
    let dec = encrypt(&enc, -3);
    println!("{} => {}", enc, dec);
}
