struct Foo {
    x: i32,
}

fn do_something(f: Foo) {
    println!("{}", f.x);
    // f is dropped here
}

fn main() {
    let mut foo = Foo { x: 42 };
    let f = &mut foo;

    // do_something(foo);
    // FAILURE: do_something(foo) would fail because
    // foo cannot be moved while mutably borrowed

    // foo.x = 13;
    // FAILURE: foo.x = 13; would fail here because
    // foo is not modifiable while mutably borrowed
    
    f.x = 13;
    // f is dropped here because it's no longer used after this point

    println!("{}", foo.x);

    // this works now because all mutable references were dropped
    foo.x = 7;

    // move foo's ownership to a function
    do_something(foo);
}
