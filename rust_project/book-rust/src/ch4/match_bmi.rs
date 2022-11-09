fn print_bmi(height: f32, weight: Option<f32>) {
    let bmi:Option<f32> = match weight {
        Some(w) => Some(w / (height / 100.0).powf(2.0)),
        None => None,
    };

    let msg = match bmi {
        Some(n) if n < 18.5 => "low fat",
        Some(n) if n < 25.0 => "normal fat",
        Some(n) if n < 30.0 => "fat lv1",
        Some(n) if n < 35.0 => "fat lv2",
        Some(n) if n < 40.0 => "fat lv3",
        Some(_) => "fat lv4",
        None => "unknown",
    };
    println!("BMI={:.1}, Judge={}", bmi.unwrap_or(0.0), msg);
}

fn main() {
    let height = 162.3;
    print_bmi(height, Some(48.0));
    print_bmi(height, Some(72.3));
    print_bmi(height, None);
}
