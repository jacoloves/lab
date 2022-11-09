fn main() {
    let height_cm = input("身長(cm)は?");
    let weight = input("体重(kg)は?");

    let height = height_cm / 100.0;
    let bmi = weight / height.powf(2.0);
    println!("BM={:.1}", bmi);

    if bmi < 18.5 { println!("low weight"); }
    else if bmi < 25.0 { println!("weight"); }
    else if bmi < 30.0 { println!("high weight lv.1"); }
    else if bmi < 35.0 { println!("high weight lv.2"); }
    else if bmi < 40.0 { println!("high weight lv.3"); }
    else { println!("high weight lv.4"); }
}

fn input(prompt: &str) -> f64 {
    println!("{}", prompt);
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).expect("input error");
    return s.trim().parse().expect("num exchange error");
}
