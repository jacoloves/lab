use std::{thread, time::Duration};

fn main() {
    let mut passed_time_sec: i32 = 0;
    let mut passed_time_hour: i32 = 0;
    let mut passed_time_min: i32 = 0;
    loop {
        progress_mark(passed_time_sec);
        progress_time(&mut passed_time_hour, &mut passed_time_min, &mut passed_time_sec);
        thread::sleep(Duration::new(1,0));
        passed_time_sec += 1;
    }
}

fn progress_time(hour: &mut i32, min: &mut i32, sec: &mut i32) {
    // hour
    *hour += *sec / 3600;
    *sec = *sec % 3600;
    
    // minute
    *min += *sec / 60;
    *sec = *sec % 60;

    eprint!("{}:{}:{}\r", hour, min, sec);
}

fn progress_mark(sec: i32) {
    let mark = vec!["-", "\\", "|", "/"];
    let pos = sec as usize % mark.len();
    eprint!("{} ", mark[pos]);
}
