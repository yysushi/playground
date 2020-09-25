fn main() {
    let mut foo = 42;
    let f = &mut foo;
    let bar = *f; // get a copy of the owner's value
    *f = 13; // set the reference's owner's value
    println!("{}", bar);
    println!("{}", foo);
    /*
     * let foo = 42;
     * let bar = foo;
     * foo = 13;
     * println!("{}", foo);
     * println!("{}", bar);
     */
}
