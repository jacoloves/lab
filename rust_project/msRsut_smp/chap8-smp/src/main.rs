struct Point<T, U> {
    x: T, 
    y: U,
}

trait Area {
    fn area(&self) -> f64;
}

struct Circle {
    radius: f64, 
}

struct Rectangle {
    width: f64, 
    height: f64,
}

impl Area for Circle {
    fn area(&self) -> f64 { 
        use std::f64::consts::PI;
        PI * self.radius.powf(2.0)
     }
}

impl Area for Rectangle {
    fn area(&self) -> f64 { 
        self.width * self.height
    }
}

#[derive(Debug, PartialEq)]
struct Point2 { 
    x: i32, 
    y: i32,
}

use std::fmt;
impl fmt::Display for Point2 {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}

trait AsJson {
    fn as_json(&self) -> String;
}

fn send_data_as_json(value: &impl AsJson) {
    println!("Sending Json data to server...");
    println!("-> {}", value.as_json());
    println!("Done!\n");
}

struct Person {
    name: String,
    age: u8,
    favorite_fruit: String,
}

struct Dog {
    name: String, 
    color: String, 
    likes_petting: bool,
}

impl AsJson for Person {
    fn as_json(&self) ->String {
        format!(
            r#"{{ "type": "person", "name": "{}", "age": {}, "favoriteFruit": "{}" }}"#,
            self.name, self.age, self.favorite_fruit
        )
    }
}

impl AsJson for Dog {
    fn as_json(&self) ->String {
        format!(
            r#"{{ "type": "dog", "name": "{}", "color": "{}", "likesPetting": "{}"}}"#,
            self.name, self.color, self.likes_petting
        )
    }
}

struct Cat {
    name: String,
    sharp_claws: bool,
}

impl AsJson for Cat {
    fn as_json(&self) -> String {
        format!(
            r#"{{ "type": "cat", "name": "{}", "sharpClaws": "{}"}}"#,
            self.name, self.sharp_claws
        )
    }
}

fn main() {
    let boolean = Point { x: true, y: false };
    let integer = Point { x: 1, y: 9 };
    let float = Point { x: 1.7, y: 4.3 };
    let string_slice = Point{ x: "high", y: "low" };

    let wont_work = Point { x: 25, y: true};

    let circle = Circle { radius: 5.0 };
    let rectangle = Rectangle { 
        width: 10.0, 
        height: 20.0,
    };

    println!("Circle area: {}", circle.area());
    println!("Rectangle area: {}", rectangle.area());

    println!("--- Unit 4/8 ---");
    let p1 = Point2 {x: 1, y: 2};
    let p2 = Point2 {x: 4, y: -3};

    if p1 == p2 {
        println!("equal!");
    } else {
        println!("not equal!");
    }

    println!("{}", p1);
    println!("{:?}", p1);
    
    println!("--- Unit 5/8 ---");
    let laura = Person {
        name: String::from("Laura"),
        age: 31,
        favorite_fruit: String::from("apples"),
    };

    let fido = Dog {
        name: String::from("Fido"),
        color: String::from("Black"),
        likes_petting: true,
    };

    send_data_as_json(&laura);
    send_data_as_json(&fido);

    let kitty = Cat {
        name: String::from("Kitty"),
        sharp_claws: false,
    };

    send_data_as_json(&kitty);
}
