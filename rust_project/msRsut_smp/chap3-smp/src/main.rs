#[derive(Debug)]
struct KeyPress(String, char);
struct MouseClick { x: i64, y: i64 }
enum WebEvent { WELoad(bool), WEClick(MouseClick), WEKeys(KeyPress) }
fn main() {
    //todo!("Display the message by using the println!() macto");
    println!("The first letter of the English alphabet is {} and the last letter is {}.", 'A', 'Z');
    let mut a_number;

    let a_word = "Ten";

    a_number = 10;

    println!("The number is {}.", a_number);
    println!("The word is {}.", a_word);

    a_number = 15;
    println!("The number is {}.", a_number);

    let shadow_num = 5;

    let shadow_num = shadow_num + 5;

    let shadow_num = shadow_num * 2;

    println!("The number is {}.", shadow_num);

    println!("----4 / 9 ----");
    let number: u32 = 14;
    println!("The number is {}.", number);

    println!("1 + 2 = {} and 8 - 5 = {} and 15 * 3 = {}", 1u32 + 2, 8i32 - 5, 15 * 3);
    println!("9 / 2 = {} but 9.0 / 2.0 = {}", 9u32 / 2, 9.0 / 2.0);

    let is_bigger = 1 > 4;
    println!("Is 1 > 4? {}", is_bigger);

    let character_1: char = 'S';
    let character_2: char = 'f';

    let smiley_face = '😃';

    let string_1 = "miley ";
    
    let string_2: &str = "ace";

    println!("{} is a {}{}{}{}.", smiley_face, character_1, string_1, character_2, string_2);

    println!("----5 / 9 ----");
    let tuple_e = ('E', 5i32, true);
    println!("Is '{}' the {}th letter of the alphabet? {}", tuple_e.0, tuple_e.1, tuple_e.2);

    struct Student { name: String, level: u8, remote: bool }

    struct Grades(char, char, char, char, f32);

    struct Unit;

    let user_1 = Student { name: String::from("Constance Sharma"), remote: true, level: 2};
    let user_2 = Student { name: String::from("Dyson Tan"), level: 5, remote: false };

    let mark_1 = Grades('A', 'A', 'B', 'A', 3.75);
    let mark_2 = Grades('B', 'A', 'A', 'C', 3.25);

    println!("{}. level {}. Remote: {}. Grades: {}, {}, {}, {}. Average: {}", user_1.name, user_1.level, user_1.remote, mark_1.0, mark_1.1, mark_1.2, mark_1.3, mark_1.4);
    println!("{}. level {}. Remote: {}. Grades: {}, {}, {}, {}. Average: {}", user_2.name, user_2.level, user_2.remote, mark_2.0, mark_2.1, mark_2.2, mark_2.3, mark_2.4);

    println!("---- 6 / 9 ----");
    
    let click = MouseClick { x: 100, y: 250 };

    let keys = KeyPress(String::from("Ctrl+"), 'N');


    println!("Mouse click location: {}, {}", click.x, click.y);

    println!("\nKeys pressed: {}{}", keys.0, keys.1);

    println!("---- 7 / 9 ----");
    goodbye();
    let formal = "Formal: Good bye.";
    let casual = "Casual: See you later!";
    goodbye2(formal);
    goodbye2(casual);
    let num = 25;
    println!("{} divided by 5 = {}", num, divide_by_5(num));
}

fn goodbye() {
    println!("Goodbye.")
}

fn goodbye2(message: &str) {
    println!("\n{}", message);
}

fn divide_by_5(num: u32) -> u32 {
    num / 5
}