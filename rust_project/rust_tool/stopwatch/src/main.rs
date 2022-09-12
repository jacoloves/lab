use std::{thread, time::Duration};

fn main() {
    let mut passed_time_sec: i32 = 0;
    let mut passed_time_hour: i32 = 0;
    let mut passed_time_min: i32 = 0;

    let judge_strs = vec!["y", "Y", "yes", "Yes", "YES"];
    let mut loop_do_flg: bool = false;
    // input command
    eprint!("start stopwatch?: ");

    let mut word = String::new();
    std::io::stdin().read_line(&mut word).ok();
    let answer = word.trim().to_string();

    for i in 0..judge_strs.len() {
        if answer == judge_strs[i] {
            loop_do_flg = true;
            break;
        }
    }

    if loop_do_flg {
        loop {
            progress_mark(passed_time_sec);
            progress_time(&mut passed_time_hour, &mut passed_time_min, &mut passed_time_sec);
            thread::sleep(Duration::new(1,0));
            passed_time_sec += 1;
        }
    }
}

fn progress_time(hour: &mut i32, min: &mut i32, sec: &mut i32) {
    // hour
    *hour += *min / 60;
    *min = *min % 60;
    
    // minute
    *min += *sec / 60;
    *sec = *sec % 60;

    // fortmat create
    let hour_str;
    let min_str;
    let sec_str;
    if *hour < 10 {
       hour_str = format!("{}{}", "0", hour);
    } else {
       hour_str = format!("{}", hour);
    }

    if *min < 10 {
        min_str = format!("{}{}", "0", min);
    } else {
        min_str = format!("{}", min);
    }

    if *sec < 10 {
        sec_str = format!("{}{}", "0", sec);
    } else {
        sec_str = format!("{}", sec);
    }

    eprint!("{}:{}:{}\r", hour_str, min_str, sec_str);
}

fn progress_mark(sec: i32) {
    let mark = vec!["-", "\\", "|", "/"];
    let pos = sec as usize % mark.len();
    eprint!("{} ", mark[pos]);
}

/*
fn in_array(needle_str: &str, haystack: Vec<str>) -> bool {
    for i in 0..haystack.len() {
        if needle_str == haystack[i] {
            return true;
        }
    }
    false
}
*/
