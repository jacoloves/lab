macro_rules! map_init {
    ($($key:expr => $val:expr),* ) => {{
        let mut tmp = std::collections::HashMap::new();
        $ (
            tmp.insert($key, $val);
        ) *
        tmp
    }}
}

fn main() {
    let week = map_init![
        "mon" => "getu",
        "tue" => "ka",
        "web" => "sui",
        "thu" => "moku",
        "fri" => "kin",
        "sat" => "do",
        "sum" => "nichi"
    ];
    println!("mon={}", week["mon"]);
    println!("web={}", week["web"]);
}