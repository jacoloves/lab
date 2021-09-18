#[derive(Debug)]

struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area4(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}

fn smp_main5() {
    let rect1 = Rectangle { width: 30, height: 50 };
    let rect2 = Rectangle { width: 10, height: 40 };
    let rect3 = Rectangle { width: 60, height: 45 };

    println!("Can rect 1 hold rect2? {}", rect1.can_hold(&rect2))
    println!("Can rect 1 hold rect3? {}", rect1.can_hold(&rect3))
}

fn smp_main4() {
    let rect1 = Rectangle { width: 30, height: 50 };

    println!(
        "The area of the rectangle is {} suare pixels.",
        rect1.area4()
    );
}

fn smp_main3() {
    let rect1 = Rectangle { width: 30, height: 50};

    println!("rect1 is {:?}", rect1);
}

fn smp_main2() {
    let rect1 = Rectangle { width: 30, height: 50};

    println!(
        "The area of the rectangle is {} square pixels.",
        area3(&rect1)
    );
}

fn area3(rectangle: &Rectangle) -> u32 {
    rectangle.width * rectangle.height
}

fn smp_main() {
    let rect1 = (30, 50);

    println!(
        "The area of the rectangle is {} square pixels.",
        area2(rect1)
        )
}

fn area2(dimensions: (u32, u32)) -> u32 {
    dimensions.0 * dimensions.1
}

fn area(width: u32, height: u32) -> u32 {
    width * height
}

fn main() {
    let width1 = 30;
    let height1 = 50;

    println!(
        "The area of the rectangle is {} square pixels.",
        area(width1, height1)
        );
    
    smp_main();
    smp_main2();
    smp_main3();
    smp_main4();
}

