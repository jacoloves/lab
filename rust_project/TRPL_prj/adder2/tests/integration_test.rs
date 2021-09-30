extern crate adder2;

mod common;

#[test]
fn it_adds_two() {
    assert_eq!(4, adder2::add_two(2));
}
