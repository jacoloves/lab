extern crate wasm_bidgen;
use wasm_bindgen::prelude::*;

mod cipher_str;

#[wasm_bindgen]
pub fn encrypt(password: &str, data:)