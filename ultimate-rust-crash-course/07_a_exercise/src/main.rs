
const STARTING_MISSILES: i32 = 8;
const READY_AMOUNT: i32 = 2;

fn main() {
    //READY_AMOUNT = 5;
    let a = 5;
    let (mut missiles): i32 = STARTING_MISSILES;
    let ready: i32 = READY_AMOUNT;
    println!("Firing {} of my {} missiles...", ready, missiles);
    println!("{} missiles left", missiles - ready);
    missiles = missiles - ready;
    println!("{} missiles left", missiles);
}
