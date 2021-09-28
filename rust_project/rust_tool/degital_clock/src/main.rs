extern crate termion;

use std::{thread, time::Duration};

fn main() {
    let mut hour = 0;
    let mut min = 0;
    let mut sec = 0;

    let mut hour_str;
    let mut min_str;
    let mut sec_str;

    loop {
        eprint!("{}{}", termion::clear::All, termion::cursor::Goto(1, 2));

        // hour translate
        if hour < 10 {
            hour_str = format!("{}{}", "0", hour);
        } else {
            hour_str = format!("{}", hour);
        }

        // minute translate
        if min < 10 {
            min_str = format!("{}{}", "0", min);
        } else {
            min_str = format!("{}", min);
        }

        // second translate
        if sec < 10 {
            sec_str = format!("{}{}", "0", sec);
        } else {
            sec_str = format!("{}", sec);
        }

        let display_clock = format!("{}:{}:{}", hour_str, min_str, sec_str);

        eprint!("{}", display_clock);

        sec += 1;
        // time calculate
        if sec == 60 {
            sec = 0;
            min += 1;
        }

        if min == 60 {
            min = 0;
            hour += 1;
        }

        thread::sleep(Duration::from_secs(1));
    }
}
