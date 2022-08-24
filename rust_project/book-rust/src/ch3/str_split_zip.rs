fn main() {
    let zipcode = "105-0011";

    println!("-- slice --");
    println!("zenhan: {}", &zipcode[..3]);
    println!("kohan: {}", &zipcode[4..]);

    println!("--- spit_at ---");
    let (zip1, zip2) = zipcode.split_at(3);
    let (zip2, zip3) = zip2.split_at(1);
    println!("zenhan: {}", zip1);
    println!("kigo: {}", zip2);
    println!("kohan: {}", zip3);

    println!("--- spit_off ---");
    let mut zip1 = String::from(zipcode);
    let mut zip2 = zip1.split_off(3);
    let zip3 = zip2.split_off(1);
    println!("zenhan: {}", zip1);
    println!("kigo: {}", zip2);
    println!("kohan: {}", zip3);

    println!("-- split --");
    let zip_a: Vec<&str> = zipcode.split('-').collect();
    println!("zenhan: {}", zip_a[0]);
    println!("kohan: {}", zip_a[1]);
}
