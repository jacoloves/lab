use std::error::Error;
use std::fs::File;
use std::io::prelude::*;

pub fn run(config: Config) -> Result<(), Box<Error>> {
    let mut f = File::open(config.filename)?;

    let mut contents = String::new();
    f.read_to_string(&mut contents)?;

    println!("With text:\n{}", contents);

    Ok(())
}

pub struct Config {
    query: String,
    filename: String,
}

impl Config {
    fn new(args: &[String]) -> Result<Config, &'static str> {
        if args.len() < 3 {
            return Err("not enough arguments");
        }

        let query = args[1].clone();
        let filename = args[2].clone();

        Ok(Config { query, filename })
    }
}

fn main() {
    pub fn search<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
        for line in contents.lines() {
            if line.contains(query) {
                result.push(line);
            }
        }

        results
    }

    #[cfg(test)]
    mod test {
        use super::*;

        #[test]
        fn one_result() {
            let query = "duct";
            let contents = "\
            Rust:
            safe, fast, productive.
            pick three.";

            assert_eq!(vec!["sage, fast, productive."], search(query, contents));
        }
    }
}
