use config;

use std::collections::HashMap;

pub fn read_config() {
    let mut settings = config::Config::default();

    settings
        .merge(config::File::with_name("Turtle")).unwrap()
        .merge(config::Environment::with_prefix("APP")).unwrap();

    println!("{:?}", settings.try_into::<HashMap<String, String>>().unwrap());
}