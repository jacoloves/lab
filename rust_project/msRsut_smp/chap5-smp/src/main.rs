use std::collections::HashMap;

fn main() {
    let mut reviews: HashMap<String, String> = HashMap::new();

    reviews.insert(String::from("Ancient Roman History"), String::from("Very accurate."));
    reviews.insert(String::from("Cooking with Rhubarb"), String::from("Sweet recipes."));
    reviews.insert(String::from("Programming in Rust"), String::from("Great examples."));

    let book: &str = "Programming in Rust";
    println!("\nReview for \'{}\': {:?}", book, reviews.get(book));

    let obsolete: &str = "Ancient Roman History";
    println!("\n'{}\' removes.", obsolete);
    reviews.remove(obsolete);

    println!("\nReview for \'{}\': {:?}", obsolete, reviews.get(obsolete));

    loop {
        println!("We loop forever");
        break;
    }

    let mut counter = 1;
    let stop_loop = loop {
        counter *= 2;
        if counter > 100 {
            break counter;
        }
    };

    println!("Break the loop at counter = {}.", stop_loop);

    let big_birds = ["ostrich", "peacock", "stork"];
    for bird in big_birds.iter() {
        println!("The {} is a big bird.", bird);
    }

    for number in 0..5 {
        println!("{}", number * 2);
    }
}
