use chrono::Local;

fn main() {
    let dt = Local::now();
    let mut hour = dt.format("%H");
    println!("{}", hour);
}
