fn main() {
  let x = 5;
  let z;
  {
    let y = 99;
    z = y;
    println!("{}, {}", x, y);
  }

  // println!("{}, {}", x, y); // Error!
  println!("{}, {}", x, z); 
}