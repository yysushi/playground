fn main() {
    let a: &'static str = "hi 🦀";
    let aa = "hi 🦀2";
    println!("{} {}", a, a.len());
    println!("{}", a == aa);
    for aaa in aa.bytes() {
        println!("{}", aaa);
    }
    let b = String::from("b");
    let c = "b";
    println!("{}", b == c);
}
