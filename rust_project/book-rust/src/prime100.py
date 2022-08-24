def is_prime(n):
    for i in range(2, n):
        if n % i == 0:
            return False

    return True

def get_primes(count):
    res = []
    i = 2
    while len(res) < count:
        if is_prime(i):
            res.append(i)
        i += 1
    return res

print(get_primes(100))
