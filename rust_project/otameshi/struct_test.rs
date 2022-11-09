struct Fstatus<'a> {
    file_names: Vec<&'a str>,
    is_files: Vec<bool>,
    byte_sizes: Vec<u64>, 
}

impl<'a> Fstatus<'a> {
    fn new() -> Self {
        Fstatus { file_names: vec![], is_files: vec![], byte_sizes: vec![] }
    }

    fn add(&mut self, file_name: &'a str, is_file: bool, byte_size: u64) -> &mut Self {
        self.file_names.push(file_name);
        self.is_files.push(is_file);
        self.byte_sizes.push(byte_size);
        self
    }
}

fn main() {
    let mut test = Fstatus::new();
    for _ in 1..=2  {
        test.add("test", true, 1234);
    }

    for i in 0..=1 {
        println!("{}: name:{} file:{} size:{}", i, test.file_names[i], test.is_files[i], test.byte_sizes[i]);
    }
}
