use libc::c_void;
use libc::open;
use libc::perror;
use libc::read;
use libc::ssize_t;
use libc::write;
use libc::O_RDONLY;
use std::env;
use std::ffi::CString;
use std::process::exit;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("{}: file name out given", args[0]);
        exit(1);
    }

    let mut cnt = 0;
    for _i in args.iter() {
        if cnt != 0 {
            do_cat(_i);
        }
        cnt = cnt + 1;
    }
    exit(0);
}

const BUFFER_SIZE: usize = 2048;

fn do_cat(_path: &str) {
    unsafe {
        let _fd: i32;
        let mut n: ssize_t;
        let buf: *mut c_void = libc::malloc(BUFFER_SIZE);

        let c_str = CString::new(_path).unwrap();
        _fd = open(c_str.as_ptr() as *const i8, O_RDONLY);

        if _fd < 0 {
            die(_path);
        }

        loop {
            n = read(_fd, buf, std::mem::size_of::<c_void>());
            if n < 0 {
                die(_path);
            }
            if n == 0 {
                break;
            }
            if write(libc::STDOUT_FILENO, buf, n as usize) < 0 {
                die(_path);
            }
        }
        if libc::close(_fd) < 0 {
            die(_path);
        }

        libc::free(buf);
    }
}

fn die(_path: &str) {
    unsafe {
        let c_str = CString::new(_path).unwrap();
        perror(c_str.as_ptr() as *const i8);
    }
    exit(1);
}
