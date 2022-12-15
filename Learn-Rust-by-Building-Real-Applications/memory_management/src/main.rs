#[derive(Debug)]
struct Test {
    t1: String,
    t2: u32
} 

impl Test {
    pub fn new(t1: String, t2: u32) -> Self {
        Self {t1, t2}
    }
}

fn main() {
    let a = 2;
    let result = stack_only(a);
    println!("{}", result);
    dbg!(result);
}

fn stack_only(b: i32) -> i32 {
    let c = 3;
    println!("{}", c);
    return b + c + stack_and_heap();
}

fn stack_and_heap() -> i32 {
    let d = 5;
    let e = Box::new(7);
    let k = e;
    let g1 = Test {
        t1: String::from("lol"), 
        t2: 32
    };
    let g2 = Test::new(String::from("lol"), 32);
    let mut l = String::from("lol");
    let lol = "lol";
    let r = &mut l;
    let a = [0; 3];
    println!("{}, {:?}, {:?}, {}, {}, {}, {}", k, g1, g2, lol, r, r.chars().nth(3).unwrap_or('f'), a[2]);
    return d + *k;
}
