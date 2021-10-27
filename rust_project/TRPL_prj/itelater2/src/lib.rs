impl Config {
    pub fn new(mut args: std::env::Args) -> Result<Config, &'static str> {
        /*
        if args.len() < 3 {
            return Err("not enough arguments");
        }
        */

        args.next();

        let query = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a query string"),
        };

        let filename = match args.next() {
            Some(arg) => arg,
            None => return Err("Didn't get a file name"),
        };

        let case_sensitive = env::var("CASE_INSENSITIVE").is_err();

        Ok(Config {
            query,
            filename,
            case_sensitive,
        })
    }

    pub fn search<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
        /*
        let mut results = Vec::new();

        for line in contents.lines() {
            if line.contains(query) {
                results.push(line);
            }
        }

        results
        */
        contents
            .lines()
            .filter(|line| line.contains(query))
            .collect()
    }
}
