use std::collections::HashMap;

fn main() {
    let tuki = ["jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"];

    let mut tuki_map: HashMap<&str, usize> = HashMap::new();

    for (i,v) in tuki.iter().enumerate() {
       tuki_map.insert(v, i+1);
    }

    println!("jun = {}月", tuki_map["jun"]);
    println!("oct = {}月", tuki_map["oct"]);
    println!("dec = {}月", tuki_map["dec"]);
}