# notes
* `use std::io`: import library, otherwise use "std::io::~~"
* `let`: create variable
* `let mut`: create mutable variable
* `String::new()`: create a new string instance
* `&`: reference of variable
* references are immutable by default
* `&mut guess`: pass mutable variable's reference
* `{}`: placeholder in "println!"
* `read_line()` returns `io:Result` type
* `io::Result` is `enumerations` variants consisting of `Ok` and `Err`
* an instance of `io::Result`has method `expect`
* If we don't call `expect`, rustc will warn
* `crate`: a package of rust code
* before write the code using external library, edit Cargo.toml
* Cargo.lock File Ensures Reproducible Builds
* if use external library, `extern crate rand;` equals to `use rand;`
* trait (not mixin) ...!!! other than module, function, struct
* Rust allows us to _shadow_ the previous value of guess with a new one
* `let guess: u32`: The colon (:) after guess tells Rust we'll annotate the variable's type.
* `match` expression is made up of arms, which consists of a pattern and the code that should be run if the value given to the beginning of the match expression fits that arm’s pattern.
* `loop`, `break`

# commands
```!bash
$ cargo new guessing_game --bin
$ cargo build # fetch external crate if necessary
# check Cargo.lock
$ cargo update # ignore Cargo.lock and overwrite
$ cargo doc --open # show doc you including external crate
$ RUST_BACKTRACE=1 ./target/debug/guessing_game
```
