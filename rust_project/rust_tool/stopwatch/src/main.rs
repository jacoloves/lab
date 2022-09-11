use std::{thread, time::Duration};

fn main() {
    let mut passed_time_sec: i32 = 0;
    let mut passed_time_hour: i32 = 0;
    let mut passed_time_min: i32 = 0;
    loop {
        // hour
        passed_time_hour += passed_time_sec / 3600;
        passed_time_sec = passed_time_sec % 3600;

        // minute
        passed_time_min += passed_time_sec / 60;
        passed_time_sec = passed_time_sec % 60;
    
        // display indication
        let s = format!("{}:{}:{}\r", passed_time_hour, passed_time_min, passed_time_sec);
        eprint!("{}", s);

        thread::sleep(Duration::new(1,0));
        passed_time_sec += 1;
    }
}
