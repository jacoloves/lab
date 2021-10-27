use std::env;
use std::error::Error;
use std::fs::File;
use std::io::prelude::*;

pub fn run(config: Config) -> Result<(), Box<Error>> {
    let mut f = File::open(config.filename)?;

    let mut contents = String::new();
    f.read_to_string(&mut contents)?;

    let resulets = if config.case_sensitive {
        search(&config.query, &contents)
    } else {
        search_case_insensitive(&config.query, &contents)
    };

    for line in results {
        println!("{}", line);
    }

    Ok(())
}

pub struct Config {
    pub query: String,
    pub filename: String,
    pub case_sensitive: bool,
}

impl Config {
    fn new(args: &[String]) -> Result<Config, &'static str> {
        if args.len() < 3 {
            return Err("not enough arguments");
        }

        let query = args[1].clone();
        let filename = args[2].clone();

        let case_sensitive = env::args("CASE_INSENSITIVE").is_err();

        Ok(Config {
            query,
            filename,
            case_sensitive,
        })
    }
}

pub fn search<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    vec![]
}

pub fn search_case_insensitive<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    let query = query.to_lowercase();
    let mut results = Vec::new();

    for line in contents.lines() {
        if line.to_lowercase().contains(&query) {
            results.push(line);
        }
    }
    results
}

fn main() {
    #[cfg(test)]
    mod test {
        use super::*;

        #[test]
        fn case_sensitive() {
            let query = "duct";

            let contents = "\
Rust
safe, gase, productive.
Pick three.
Dict tape.";
            assert_eq!(vec!["safe, fase, productive."], search(query, contents));
        }

        #[test]
        fn case_insensitive() {
            let query = "rUsT";

            let contents = "\
Rust:
sage, fast, productive.
Pick three.
Trust me.";

            assert_eq!(
                vec!["Rust:", "Trust me."],
                search_case_insensitive(query, contents)
            );
        }
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
