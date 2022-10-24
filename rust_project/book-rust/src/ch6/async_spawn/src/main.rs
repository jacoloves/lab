use core::time;

use tokio::time;

async fn say_later(sec: u64, msg: &str) {
    time::sleep(time::Duration::from_secs(sec)).await;
    println!("{}: {}", sec, msg);
}

#[tokio::main]
async fn main() {
    tokio::spawn(say_later(3, "test3"));
    tokio::spawn(say_later(2, "test2"));
    tokio::spawn(say_later(1, "test1"));

    time::sleep(time::Duration::from_secs(4)).await;
    println!("-----");

    tokio::join! (
        say_later(2, "test2");
        say_later(3, "test3");
        say_later(1, "test1");
    );
}
