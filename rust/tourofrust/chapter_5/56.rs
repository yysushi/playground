// struct Foo {
//     i: &i32,
// }
struct Foo<'a> {
    i: &'a i32,
}

fn main() {
    let x = 42;
    let foo = Foo { i: &x };
    println!("{}", foo.i);
}
