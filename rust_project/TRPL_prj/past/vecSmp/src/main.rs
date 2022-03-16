#![allow(unused)]
fn vec_smp3() {
    enum SpreadsheetCell {
        Int(i32),
        Float(f64),
        Text(String),
    }

    let row = vec![
        SpreadsheetCell::Int(3),
        SpreadsheetCell::Text(String::from("blue")),
        SpreadsheetCell::Float(10.12),
    ];
}

fn vec_smp2() {
    let mut v = vec![100, 32, 57];
    for i in &mut v {
        *i += 50;
        println!("{}", i);
    }
}
fn vec_smp() {
    let v = vec![100, 32, 57];
    for i in &v {
        println!("{}", i);
    }
}
fn main() {
    //let v: Vec<i32> = Vec::new();
    //    let v = vec![1, 2, 3];
    //    let mut v = Vec::new();
    //
    //    v.push(5);
    //    v.push(6);
    //    v.push(7);
    //    v.push(8);
    //
    //  let v = vec![1, 2, 3, 4, 5];

    //    let third: &i32 = &v[2];
    //    let third: Option<&i32> = v.get(2);

    // let first = &v[0];
    // v.push(6);
    // println!("The first element is: {}", first);
    //
    vec_smp();
    vec_smp2();
}
