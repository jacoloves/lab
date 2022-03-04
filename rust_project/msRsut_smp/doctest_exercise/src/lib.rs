/// div test
/// ```
/// let res = doctest_exercise::div(10, 2);
/// assert_eq!(res, 5);
/// ```
/// 
/// div test2
/// ```
/// let res = doctest_exercise::div(6, 2);
/// assert_eq!(res, 3);
/// ```
/// 
/// # panic tests
/// ```
/// let res = doctest_exercise::div(3, 0);
/// ```

pub fn div(a: i32, b: i32) -> i32 {
    if b == 0 {
        panic!("Diveide-by-zero error");
    }
    a / b
}

/// # examples
/// # subtest
/// ```
/// let res = doctest_exercise::sub(9, 2);
/// assert_eq!(res, 7);
/// ```
/// 
/// # subtest2
/// ```
/// let res = doctest_exercise::sub(6, 9);
/// assert_eq!(res, -3);
/// ```

pub fn sub(a: i32, b: i32) -> i32 {
    a - b
}
