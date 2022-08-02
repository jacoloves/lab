fn main() {
    let s = "苦しむ人はどの日も悪い日である。";

    let s2 = s.replace("苦しむ人", "陽気な人");
    let s3 = s2.replace("悪い日", "宴会");

    println!("置換前: {}\n置換後: {}", s, s3);
}
