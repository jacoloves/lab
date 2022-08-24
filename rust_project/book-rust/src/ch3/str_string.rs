fn main() {
    let ss: &str = "気前よく与えるなら豊かになる";
    let so1: String = String::from(ss);
    let so2: String = ss.to_string();

    let ss2: &str = &so1;
    let ss3: &str = so1.as_str();

    println!("{}\n{}\n{}\n{}", so1, so2, ss2, ss3);

    println!("{:p}\n{:p}", ss2, ss3);
}
