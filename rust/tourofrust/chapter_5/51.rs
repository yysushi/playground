struct Foo {
    x: i32,
}

fn do_something(f: &mut Foo) {
    f.x += 1;
    // mutable reference f is dropped here
}

fn main() {
    // let mut foo = Foo { x: 42 };
    // do_something(&mut foo);
    // println!("{}", foo.x);
    // // because all mutable references are dropped within
    // // the function do_something, we can create another.
    // do_something(&mut foo);
    // println!("{}", foo.x);
    // // foo is dropped here

    let mut foo = Foo { x: 42 };
    let y = &mut foo;
    do_something(y);
    // do_something(&mut foo);
    do_something(y);
    do_something(&mut foo);
    // println!("{}", foo.x);
}
