struct Person {
    first: String,
    middle: Option<String>,
    last: String,
}

fn build_full_name(person: &Person) -> String {
    let mut full_name = String::new();
    full_name.push_str(&person.first);
    full_name.push_str(" ");
    
    full_name.push_str(&person.last);
    full_name
}

fn chap6_person_main() {
    let john - Person { first: "John", last: ""}
}

fn main() {
    //panic!("Farewell!");

    let fruits = vec!["banana", "apple", "coconut", "orange", "strawberry"];

    let first = fruits.get(0);
    println!("{:?}", first);

    let third = fruits.get(2);
    println!("{:?}", third);

    let non_existent = fruits.get(99);
    println!("{:?}", non_existent);

    for &index in [0, 2, 99].iter() {
        match fruits.get(index) {
            Some(&"coconut") => println!("Coconuts are awesome!!!"),
            Some(fruit_name) => println!("It's a delicious {}!", fruit_name),
            None => println!("There is no fruit! :("),
        }
    }

    let a_number: Option<u8> = Some(7);
    if let Some(7) = a_number {
        println!("That's my lucky number!");
    }

    // let gift = Some("cnady");
    // assert_eq!(gift.unwrap(), "candy");

    // let empty_gift: Option<&str> = None;
    // assert_eq!(empty_gift.unwrap(), "candy");

    // let a = Some("value");
    // assert_eq!(a.expect("fruits are healthy"), "value");

    // let b: Option<&str> = None;
    // b.expect("fruits are healthy");

    assert_eq!(Some("dog").unwrap_or("cat"), "dog");
    assert_eq!(None.unwrap_or("cat"), "cat");
}
