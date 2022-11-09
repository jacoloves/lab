fn main() {
    let age = 8;
    let age_str = match age {
        0 => "nyuji",
        1..=5=> "yoji",
        6..=11 => "kid",
        _ => "man",
    };

    println!("{}歳は{}料金", age, age_str);
}
