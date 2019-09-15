mod colors;
mod commands;
mod prompt;
mod shell;
mod config;

fn main() {
    config::read_config();
    shell::run();
}