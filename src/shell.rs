use std::io::stdin;
use std::process::Command;

use crate::commands;
use crate::prompt;

pub fn run() {
    loop {
        prompt::print();

        let mut input = String::new();
        stdin().read_line(&mut input).unwrap();

        let mut parts = input.trim().split_whitespace();
        let command = parts.next().unwrap();
        let args = parts;

        match command {
            "cd" => commands::change_dir(args),
            "exit" => return,
            command => {
                let child = Command::new(command).args(args).spawn();

                match child {
                    Ok(mut child) => {
                        child.wait().expect("Error");
                    }
                    Err(e) => eprintln!("{}", e),
                };
            }
        }
    }
}
