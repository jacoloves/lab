fn echo(s: &'static str) {
    println!("{}", s);
}

fn main() {
    echo("hello");
    echo("world");


}
