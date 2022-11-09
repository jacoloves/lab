use std::sync::mpsc;
use std::thread;
use std::time::Duration;

fn sleep_sender(name: &str, sender: mpsc::Sender<String>) {
    for i in 1..=5 {
        let msg = format!("{}: {}", name, i);
        sender.send(msg).unwrap();
        thread::sleep(Duration::from_millis(1000));
    }
    sender.send("quit".to_string()).unwrap();
}

fn main() {
    let (tx, rx) = mpsc::channel::<String>();

    let sender = tx.clone();
    thread::spawn(|| {
       sleep_sender("taro", sender) 
    });

    let sender = tx.clone();
    thread::spawn(|| {
        sleep_sender("jiro", sender)
    });

    loop {
        let buf = rx.recv().unwrap();
        println!("[reception]{}", buf);
        if buf == "quit" { break; }
    }
}
