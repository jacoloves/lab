def encrypt(text, shift):
    a = ord('A')
    conv = lambda n: chr((ord(n) - a + shift) % 26 + a)
    encl = lambda n: conv(n) if 'A' <= n <= 'Z' else n
    return ''.join([encl(n) for n in text])

enc = encrypt("I LOVE YOU.", 3)
dec = encrypt(enc, -3)
print(enc, "=>", dec)
