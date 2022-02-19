use std::fs::File;
use std::io::{Error, Read};
use std::path::PathBuf;

fn read_file_contents(path: PathBuf) -> Result<String, Error> {
    let mut string = String::new();

    let mut file: File = match File::open(path) {
        Ok(file_handle) => file_handle, 
        Err(io_error) => return Err(io_error), 
    };

    match file.read_to_string(&mut string) {
        Ok(_) => (), 
        Err(io_error) => return Err(io_error),
    };

    Ok(string)
}

fn main() {
    if read_file_contents(PathBuf::from("src/main.rs")).is_ok() {
        println!("The program found the main file.");
    }
    if read_file_contents(PathBuf::from("non-existent-file.txt")).is_err() {
        println!("The program reported an error for the file thar doesn't exitst.");
    }
}
