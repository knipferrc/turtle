use std::env;
use std::io::{stdout, Write};

use whoami;

use crate::colors;

pub fn print() {
    let path = env::current_dir().unwrap();
    let split_path: Vec<&str> = path.to_str().unwrap().split("/").collect();

    let prompt = format!(
        "{}{}{}{}{}{}{}{}{}{}{}{}{} ",
        colors::GREEN,
        colors::ANSI_BOLD,
        whoami::username(),
        "@",
        whoami::hostname(),
        ":",
        colors::ANSI_BOLD,
        colors::ANSI_COLOR_MAGENTA,
        "~/",
        split_path[split_path.len() - 2],
        "/",
        split_path[split_path.len() - 1],
        colors::RESET
    );

    print!("{}", prompt);
    stdout().flush().ok().expect("Could not flush stdout");
}
