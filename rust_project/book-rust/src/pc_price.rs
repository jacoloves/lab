fn main() {
    let pc_price = 98000.0;
    let a_price = pc_price - (pc_price * 0.2) + 1200.0;
    let b_price = pc_price - (pc_price * 0.1);

    if a_price < b_price {
        println!("A");
    } else if a_price > b_price {
        println!("B");
    } else {
        println!("same");
    }
}
