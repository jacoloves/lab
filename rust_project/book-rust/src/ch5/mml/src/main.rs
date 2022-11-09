mod mml_parser;
mod wav_writer;

fn main() {
    let mml = format!("{}{}",
        "o5 l4 cdef edl2c l4 efga gfl2e",
        "l4 crcr crcr l8 ccdd eeff l4 ed l2c");

    let notes = mml_parser::parse(mml);
    let bpm = 120.0;
    wav_writer::write("kaeru.wav", notes, bpm);

    let mml = format!("{}{}{}",
        "o5 l4 ccgg aal2g l4 ffee ddl2c",
        "l4 ggff eel2d l4 ggff eel2d",
        "l4 ccgg aal2g l4 ffee ddl2c"
    );
    let notes = mml_parser::parse(mml);
    let bpm = 120.0;
    wav_writer::write("kirakira.wav", notes, bpm);
}
