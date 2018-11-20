use std::env;
use std::io::{stdin, stdout, Write};
use std::process::Command;

fn main() {
    loop {
        let path = env::current_dir().unwrap();

        print!("{} 🐢  >= ", path.display());
        stdout().flush().expect("Error flushing output");

        let mut input = String::new();
        stdin().read_line(&mut input).unwrap();

        let mut parts = input.trim().split_whitespace();
        let command = parts.next().unwrap();
        let args = parts;

       let mut child = if cfg!(target_os = "windows") {
            Command::new(command)
                .args(args)
                .spawn()
                .unwrap()
        } else {
            Command::new(command)
                .args(args)
                .spawn()
                .unwrap()
        };
        
        child.wait().expect("Failed executing command");
    }
}
