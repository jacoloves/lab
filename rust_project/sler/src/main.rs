fn clear() {
   println!("\x1b[2J");
   println!("\x1b[H");
}

fn print_sler() {
    println!("
        ________________      _____ 
       /   ____________/     /    / 
      /    /                /    / 
     /    /___________     /    / 
    /____________    /    /    /              _____      _______
                /   /    /    /              / __ \\     /  ____/ 
    ___________/   /    /    /___________   / ____\\    /  / 
   /              /    /                /  / /____    /  /
  ----------------    /________________/   \\_____/   /__/
      ");
}

fn main() {
    self::clear();
    print_sler();
}