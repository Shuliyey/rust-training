#![allow(dead_code)]

fn main() {
  let bunnies: i32 = 4;
  let (bunnies, carrots) = (8, 50);
  // bunnies = 5; error
  println!("bunnies: {}, carrots: {}", bunnies, carrots);

  let mut bunnies = 5;
  bunnies = 10;
  println!("mut bunnies: {}", bunnies);
}