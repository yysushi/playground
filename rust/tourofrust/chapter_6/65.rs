fn main() {
    let a = "hi 🦀";
    println!("{}", a.len());
    let first_word = &a[0..2];
    let second_word = &a[3..7];
    let space = &a[2..3];
    // let half_crab = &a[3..5]; FAILS
    // Rust does not accept slices of invalid unicode characters
    println!("{} {}", first_word, second_word);
    println!("{}!", space);
    let b = "hi 😊🦀";
    match b.find("🦀") {
        Some(i) => println!("{}", i),
        None => {}
    };
}
