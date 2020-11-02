struct SeaCreature {
    noise: String,
}

impl SeaCreature {
    fn get_sound(&self) -> &str{
        &self.noise
    }
}

trait NoiseMaker {
    fn make_noise(&self);
}

trait LoudNoiseMaker: NoiseMaker {
    fn make_a_lot_of_noises(&self) {
        self.make_noise();
        self.make_noise();
        self.make_noise();
    }
}

impl NoiseMaker for SeaCreature {
    fn make_noise(&self) {
        println!("{}", self.get_sound());
    }
}

impl LoudNoiseMaker for SeaCreature {}

fn main() {
    let creature = SeaCreature {
        noise: String::from("blub"),
    };
    // creature.make_noise();
    creature.make_a_lot_of_noises();
}
