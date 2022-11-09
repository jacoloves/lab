struct Body {
    weight: f64,
    height: f64,
}

fn main() {
    let ichiro = Body{weight: 80.0, height: 165.0};
    let jiro = Body{weight: 65.0, height: 170.0};

    println!("Ichiro={:.1}", calc_bmi(&ichiro));
    println!("Jiro={:.1}", calc_bmi(&jiro));
}

fn calc_bmi(body: &Body) -> f64 {
    let h = body.height / 100.0;
    body.weight / h.powf(2.0)
}
