fn smp_num() {
    let x = 5;
    let y = x;

    println!("x = {}, y = {}", x, y);
}

fn smp_copy() {
    let s1 = String::from("hello");
    let s2 = s1.clone();

    println!("s1 = {}, s2 = {}", s1, s2);
}

fn main() {
    let mut s = String::from("hello");

    s.push_str(", world");

    println!("{}", s);
    
    smp_copy();
}
