fn main() {
    let year = 2023;
    let month = 12;
    let day = 1;
    println!("JP:{0}/{1}/{2}", year, month, day);
    println!("UP:{1}/{2}/{0}", year, month, day);
    println!("UK:{2}/{1}/{0}", year, month, day);
}
