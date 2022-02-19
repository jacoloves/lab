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
}
