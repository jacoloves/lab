function isBook(inf: Book | Magazine): inf is Book {
    return (inf as Book).isbn != undefined;
}

let i = getInfo();
if (isBook(i)) {
    console.log(i.isbn);
} else {
    console.log(i.mcode);
}

