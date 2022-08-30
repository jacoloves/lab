enum Coin {
    Coin1(isize),
    Coin5(isize),
    Coin10(isize),
    Coin50(isize),
    Coin100(isize),
    Coin500(isize),
}

impl Coin {
    fn calc_price(&self) -> isize {
        match *self {
            Coin::Coin1(v) => v,
            Coin::Coin5(v) => v * 5,
            Coin::Coin10(v) => v * 10,
            Coin::Coin50(v) => v * 50,
            Coin::Coin100(v) => v * 100,
            Coin::Coin500(v) => v * 500,
        }
    }
}

fn main() {
    let wallet: Vec<Coin> = vec![
        Coin::Coin1(3),
        Coin::Coin5(10),
        Coin::Coin10(5),
        Coin::Coin50(1),
        Coin::Coin100(1),
        Coin::Coin500(5),
    ];

    let total = wallet.iter()
        .fold(0, |sum, v| sum + v.calc_price());
    println!("財布の合計は{}円です。", total);
}
