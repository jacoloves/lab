use std::fs::File;
use std::io;
use std::io::ErrorKind;
use std::io::Read;

fn panic_smp() {
    pub struct Guess {
        value: u32,
    }

    impl Guess {
        pub fn new(value: u32) -> Guess {
            if value < 1 || value > 100 {
                panic!("Guess value must be between 1 and 100, got {}.", value);
            }

            Guess { value }
        }

        pub fn value(&self) -> u32 {
            self.value
        }
    }
}

fn smp_main() {
    let f = File::open("hello2.txt")?;
}

fn read_username_from_file3() -> Result<String, io::Error> {
    let mut s = String::new();

    File::open("hello.txt")?.read_to_string(&mut s)?;

    Ok(s)
}

fn read_username_from_file2() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt");
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("hello2.txt");

    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();

    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}

fn expect_smp() {
    let f = File::open("hello2.txt").expect("Failed to open hello2.txt");
}

fn unwrap_smp() {
    let f = File::open("hello2.txt").unwrap();
}

fn file_open_error() {
    let f = File::open("hello.txt");

    let f = match f {
        Ok(file) => file,
        Err(ref error) if error.kind() == ErrorKind::NotFound => match File::create("hello.txt") {
            Ok(fc) => fc,
            Err(e) => {
                panic!("Tried to create file but there was a problem: {:?}", e)
            }
        },
        Err(error) => {
            panic!("There was a problem opening the file: {:?}", error)
        }
    };
}

fn main() {
    //    panic!("crash and burn");
    //let v = vec![1, 2, 3];
    //v[99];
    //    let f: u32 = File::open("hello.txt");
    //let f = File::open("hello.txt");

    //let f = match f {
    //    Ok(file) => file,
    //    Err(error) => {
    //        panic!("There was a problem opening the file: {:?}", error)
    //    }
    //};

    //file_open_error();
    //    unwrap_smp();
    //expect_smp();
    smp_main();
}
