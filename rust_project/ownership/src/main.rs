fn smp_main6() {
    let mut s = String::from("hello world");

    let word = first_word(&s);

    s.clear();
}

fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}

fn smp_main5() {
    let mut s = String::from("hello");

    change(&mut s);

    println!("{}", s)
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}

fn smp_main4() {
    let s1 = String::from("hello");

    let len = calculate_length2(&s1);

    println!("The length of '{}' is {}.", s1, len);

}

fn calculate_length2(s: &String) -> usize {
    s.len()
} 

fn smp_main3() {
    let s1 = String::from("hello");

    let (s2, len) = calculate_length(s1);

    println!("The length of '{}' is {}.", s2, len);
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len();

    (s, length)
}

fn smp_main2() {
    let s1 = gives_ownership();

    let s2 = String::from("hello");

    let s3 = takes_and_gives_back(s2);
}

fn takes_and_gives_back(a_string: String) -> String {
    a_string
}

fn gives_ownership() -> String {
    let some_string = String::from("hello");

    some_string
}

fn makes_copy(some_integer: i32) {
    println!("{}", some_integer);
}

fn takes_ownership(some_string: String) {
    println!("{}", some_string);
}

fn smp_main() {
    let s = String::from("hello");

    takes_ownership(s);

    let x = 5;

    makes_copy(x);
}

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

    smp_num();

    smp_main();
    smp_main2();
    smp_main3();
    smp_main4();
    smp_main5();
    smp_main6();
}
