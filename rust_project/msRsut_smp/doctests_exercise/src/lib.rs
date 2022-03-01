pub fn div(a: i32, b: i32) -> i32 {
    if b== 0 {
        panic!("Divide-by-zero error");
    }
    a / b
}

pub fn sub(a: i32, b: i32) -> i32 {
    a - b
}