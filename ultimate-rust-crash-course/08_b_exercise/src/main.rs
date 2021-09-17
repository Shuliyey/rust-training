fn main() {
    println!("Hello, world!");
    println!("do_stuff({}, {}) = {}", 20.0, 39.97, do_stuff(20.0, 39.97));
}

fn do_stuff(qty: f64, oz: f64) -> f64 {
    println!("{} {}-oz sarsaparilla(s)!", qty, oz);
    qty * oz
}