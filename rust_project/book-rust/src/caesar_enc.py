def encrypt(text, shift):
    code_a = ord('A')
    code_z = ord('Z')

    result = ""
    for ch in text:
        code = ord(ch)
        if code_a <= code <= code_z:
            code = (code - code_a + shift) % 26 + code_a
        result += chr(code)

    return result

enc = encrypt("I LOVE YOU.", 3)
dec = encrypt(enc, -3)
print(enc, "=>", dec)
