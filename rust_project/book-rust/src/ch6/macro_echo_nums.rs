#[macro_export]
macro_rules! echo_nums {
    ($($num:expr), *) => {
        $(
            print!("{}, ", $num);
        ) *
        println!("");
    }
}

fn main() {
    echo_nums![10, 20, 30, 40, 50];
}