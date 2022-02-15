#[derive(PartialEq, Debug)]
struct Car {
    color: String, 
    motor: Transmission, 
    roof: bool,
    age: (Age, u32),
}

#[derive(PartialEq, Debug)]
enum Transmission {Manual, SemiAuto, Automatic}

#[derive(PartialEq, Debug)]
enum Age {New, Used}

fn car_factory(color: String, motor: Transmission, roof: bool, miles: u32) -> Car{
    Car {
        color: color,
        motor: motor,
        roof: roof,
        age: car_quality(miles),
    }
}

fn car_quality(miles: u32) -> (Age, u32) {
    let quality: (Age, u32) = (Age::New, 0);
    quality
}

fn main() {
    println!("--- unit 3/7 ---");

    let three_nums = vec![15, 3, 46];
    println!("Initial vector: {:?}", three_nums);

    let zeroes = vec![0; 5];
    println!("Zeroes: {:?}", zeroes);

    let mut fruit = Vec::new();
    fruit.push("Apple");
    fruit.push("Banana");
    fruit.push("Cherry");
    println!("Fruits: {:?}", fruit);

    println!("Pop off: {:?}", fruit.pop());
    println!("Fruits: {:?}", fruit);

    let mut index_vec = vec![15, 3, 46];
    let three = index_vec[1];
    println!("Vector: {:?}, three = {}", index_vec, three);
    index_vec[1] = index_vec[1] + 5;
    println!("Vector: {:?}", index_vec);

    println!("--- unit 4/7 ---");
    let colors = ["Blue", "Green", "Red", "Silver"];

    let mut car: Car;
    let mut engine: Transmission = Transmission::Manual;

    car = car_factory(String::from(colors[0]), engine, true, 0);
    println!("Car order 1: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);

    engine = Transmission::SemiAuto;
    car = car_factory(String::from(colors[1]), engine, false, 100);
    println!("Car order 2: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);

    engine = Transmission::Automatic;
    car = car_factory(String::from(colors[2]), engine, true, 200);
    println!("Car order 3: {:?}, Hard top = {}, {:?}, {}, {} miles", car.age.0, car.roof, car.motor, car.color, car.age.1);

    println!("--- unit 5/7 ---");
    if 1 == 2 {
        println!("True, the numbers are equal.");
    } else {
        println!("False, the numbers are not equal.");
    }

    let formal = true;
    let greeting = if formal { 
        "Good day to you." 
    } else {
        "Hey!"
    };
    println!("{}", greeting);

    let num = 500;
    let out_of_range: bool;
    if num < 0 {
        out_of_range = true;
    } else if num == 0 {
        out_of_range = true;
    } else if num > 512 {
        out_of_range = true;
    } else {
        out_of_range = false;
    }

    println!("{}", out_of_range);
}
