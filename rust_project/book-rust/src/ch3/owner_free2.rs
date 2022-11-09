fn main() {
    {
        let s1 = String::from("真実はワインの中にある");
        let s3 = String::from("ブドウ畑と美人は手がかかる");
        {
            let s2 =  s1;
            println!("{}", s2);
        }
        println!("{}", s3);
    }
}
