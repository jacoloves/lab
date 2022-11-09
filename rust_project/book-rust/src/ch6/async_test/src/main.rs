#[tokio::main]
async fn main() {
    let f = say_later("test");

    println!("ttttes");

    f.await;
}

async fn say_later(msg: &'static str) {
    println!("{}", msg);
}