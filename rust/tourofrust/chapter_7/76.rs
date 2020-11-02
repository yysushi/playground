struct SeaCreatture {
    noise: String,
}

impl SeaCreatture {
    fn get_sound(&self) -> &str {
        &self.noise
    }
}

fn main() {
    let creature = SeaCreatture {
        noise: String::from("blub"),
    };
    println!("{}", creature.get_sound())
}
