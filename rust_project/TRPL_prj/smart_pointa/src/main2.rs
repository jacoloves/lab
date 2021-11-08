struct CustromSmartPointer {
    data: String,
}

impl Drop for CustromSmartPointer {
    fn drop(&mut self) {
        println!("Dropping CustromSmartPointer with data `{}`!", self.data);
    }
}

fn main() {
    let c = CustromSmartPointer {
        data: String::from("my stuff"),
    };
    let d = CustromSmartPointer {
        data: String::from("other stuff"),
    };
    println!("CustromSmartPointers created.");
    drop(c);
    println!("CustromSmartPointer dropped before the end of main.");
}
