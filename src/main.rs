use std::env;
use std::str::{SplitWhitespace};
use std::io::{stdin, stdout, Write};
use std::path::Path;
use std::process::{Command};

mod colors;

fn change_dir(args: SplitWhitespace) {
    let new_dir = args.peekable().peek().map_or("/", |x| *x);
    let root = Path::new(new_dir);
    if let Err(e) = env::set_current_dir(&root) {
        eprintln!("{}", e);
    }
}

fn generate_prompt() -> String {
    let path = env::current_dir().unwrap();
    let split_path: Vec<&str> = path.to_str().unwrap().split("/").collect();
    
    let prompt = format!(
        "{}{}{}{}{}{}{} ", 
        colors::ANSI_BOLD,
        colors::GREEN,
        "~/",
        split_path[split_path.len() - 2],
        "/",
        split_path[split_path.len() - 1],
        colors::RESET
    );

    return prompt;
}

fn run() {
    loop {
        let prompt = generate_prompt();
        print!("{}", prompt);
        stdout().flush().ok().expect("Could not flush stdout");

        let mut input = String::new();
        stdin().read_line(&mut input).unwrap();
        
        let mut parts = input.trim().split_whitespace();
        let command = parts.next().unwrap();
        let args = parts;

        match command {
            "cd" => change_dir(args),
            "exit" => return,
            command => {
                let child = Command::new(command)
                    .args(args)
                    .spawn();

                match child {
                    Ok(mut child) => { child.wait().expect("Error"); },
                    Err(e) => eprintln!("{}", e),
                };
            }
        }
    }
}

fn main() {
    run();
}