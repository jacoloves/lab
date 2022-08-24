struct Item(String, i64);
fn main() {
    let banana = Item("banana".to_string(), 300);
    let apple = Item("apple".to_string(), 200);
    let mango = Item("mango".to_string(), 500);

    let items = vec![banana, apple, mango];
    let total = print_add_sum_items(&items);
    println!("sum {}", total);
}

fn print_tuple(item: &Item) {
    println!("buy {} for ${}", item.0, item.1);
}

fn print_add_sum_items(items: &Vec<Item>) -> i64 {
    let mut total = 0;
    for it in items {
       print_tuple(&it) ;
       total += it.1;
    }
    total
}
