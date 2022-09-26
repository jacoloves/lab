use std::net::{TcpListener, TcpStream};
use std::io::{BufReader, BufRead, Write};
use std::thread;
use std::sync::mpsc;
use std::time::Duration;

fn main() {
    let server_addr= "127.0.0.1:8888";
    let (tx, rx) = mspc::channel::<String>();
    let mut clients: Vec<TcpStream> = Vec::new();

    let server = TcpListener::bind(server_addr)
        .expect("server not boot");
    server.set_nonblocking(true).expect("not use");
    println!("{} server boot", server_addr);

    loop {
        if let Ok((client, addr)) = server.accept() {
            println!("client connect: {}", addr);
            clients.push(client.try_clone().unwrap());
            start_thread(client, tx.clone());
        }

        if let Ok(msg) = rx.try_recv() {
            clients = send_all(clients, &msg);
        }
        thread::sleep(Duration::from_millis(100));
    }
}

fn start_thread(client: TcpStream, tx: mpsc::Sender<String>) {
    let mut reader = BufReader::new(client);
    thread::spawn (move || loop {
        let mut msg = String::new();
        if let Ok(n) = reader.read_line(&mut msg) {
            if n > 0 { tx.send(msg).unwrap(); }
        }
        thread::sleep(Duration::from_millis(100));
    });
}

fn send_all(clients: Vec<TcpStram>, s: &str) -> Vec<TcpStream> {
    let mut collector = vec![];
    for mut socket in clinets.into_iter() {
        let bytes = String::from(s).into_bytes();
        if let Err(e) = socket.write_all(&bytes) {
            println!("send error: {}", e);
            continue;
        }
        collector.push(socket);
    }
    collector
}
