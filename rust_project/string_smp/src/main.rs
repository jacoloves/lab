fn str_smp() {
    let mut s = String::new();

    let data = "initial contents";

    let s2 = data.to_string();

    let s2 = "initial contents".to_string();

    let mut s3 = String::from("foo");
    s3.push_str("bar");

    let mut s5 = String::from("foo");
    let s4 = "bar";
    s5.push_str(s4);
    println!("s4 is {}", s4);

    let s6 = String::from("hello, ");
    let s7 = String::from("world!");
    let s8 = s6 + &s7;
    println!("{}", s8);

    let s9 = String::from("tic");
    let s10 = String::from("tac");
    let s11 = String::from("toe");

    let ss = format!("{}-{}-{}", s9, s10, s11);

    println!("{}", ss);

    let hello = "teagaf";
    let ans = &hello[0..4];

    println!("{}", ans);

    for c in "3dsfjadlklk".chars() {
        println!("{}", c);
    }

    for b in "3dsfjadlklk".bytes() {
        println!("{}", b);
    }
}

fn main() {
    str_smp();
}
