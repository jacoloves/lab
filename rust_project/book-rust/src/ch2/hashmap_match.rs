use std::collections::HashMap;

fn main() {
    let mut map = HashMap::new();
    map.insert("A", 30);
    map.insert("B", 50);

    match map.get("D") {
        None => println!("Dは存在しない"),
        Some(v) => println!("D={}", v),
    }
}
