use std::env;
use std::process::exit;
use libc::open;
use libc::O_RDONLY;
use libc::perror;
use std::ffi::{CString, CStr};
use std::os::raw::c_char;
use syscall::call;

pub const BUFFER_SIZE: usize = 2048;

fn main() {
    //let mut _i :u32;
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("{}: file name out given", args[0]);
        exit(1);
    }

    for _i in args.iter() {
        do_cat(_i);
    }
    exit(0)
}

fn do_cat(_path: &str) {
    let mut fd :Result<usize>;
    let mut buf :String;

    fd = open(_path, O_RDONLY);
    if fd < 0 {
        die(_path);
    }
}

fn die(s: &str) {
    perror(s);
    exit(1);
}
