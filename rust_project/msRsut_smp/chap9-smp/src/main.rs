use regex::Regex;

mod authentication;

fn main() {
    /*
    let user = authentication::User::new("jeremy", "super-secret");

    println!("The username is: {}", user.username);
    user.set_password("even-more-secret");
    */
    let re = Regex::new(r"^\d{4}-\d{2}-\d{2}$").unwrap();
    println!("Did our date match? {}", re.is_match("2014-01-01"));
}
