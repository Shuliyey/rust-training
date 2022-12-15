use hello::greet;
use rand::Rng;
use rand::thread_rng;

fn main() {
    greet();
    println!("Hello, world!");
    println!("random number: {}", thread_rng().gen_range(0, 100));
}
