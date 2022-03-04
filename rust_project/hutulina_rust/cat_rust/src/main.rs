use std::env;
use std::ffi::CString;
use std::process::exit;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        do_stdout();
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
        let mut n: libc::ssize_t;
        let buf: *mut libc::c_void = libc::malloc(BUFFER_SIZE);

        let c_str = CString::new(_path).unwrap();
        _fd = libc::open(c_str.as_ptr() as *const i8, libc::O_RDONLY);

        if _fd < 0 {
            die(_path);
        }

        loop {
            n = libc::read(_fd, buf, std::mem::size_of::<libc::c_void>());
            if n < 0 {
                die(_path);
            }
            if n == 0 {
                break;
            }
            if libc::write(libc::STDOUT_FILENO, buf, n as usize) < 0 {
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
        libc::perror(c_str.as_ptr() as *const i8);
    }
    exit(1);
}

fn do_stdout() {
    unsafe {
        let buf: *mut libc::c_void = libc::malloc(BUFFER_SIZE);
        let mut n: libc::ssize_t;
        loop {
            n = libc::read(
                libc::STDOUT_FILENO,
                buf,
                std::mem::size_of::<libc::c_void>(),
            );
            if libc::write(libc::STDOUT_FILENO, buf, n as usize) < 0 {
                panic!("write_error");
            }
        }
    }
}
