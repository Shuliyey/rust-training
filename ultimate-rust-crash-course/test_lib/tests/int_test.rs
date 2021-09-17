use test_lib;

mod common;

#[test]
fn it_adds_two() {
    common::setup();
    assert_eq!(4, test_lib::add_two(2));
}