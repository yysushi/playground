fn main() {
    let mut x = 1;
    loop {
        x += 1;
        if x == 42 {
            break;
        }
    }
    println!("x is {}", x);
}
