use std::{thread, time};

fn sleep_print(name: &str) {
    for i in 1..=3 {
        println!("{}: i={}", name, i);
        thread::sleep(time::Duration::from_millis(1000));
    }
}


fn main() {
    println!("--- no thread ---");
    sleep_print("no thread");

    println!("--- use thread ---");

    thread::spawn(|| {
        sleep_print("jiro")
    });

    thread::spawn(|| {
        sleep_print("subro")
    });

    sleep_print("taro");
}
