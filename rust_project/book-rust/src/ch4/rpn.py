stack = []

s = input('RPN:')

tokens = s.split()
for t in tokens:
    t = t.strip()
    try:
        stack.append(float(t))
    except ValueError:
        b = stack.pop()
        a = stack.pop()
        if t == "+": stack.append(a + b)
        elif t == "-": stack.append(a - b)
        elif t == "*": stack.append(a * b)
        elif t == "/": stack.append(a / b)
        else: raise Exception("undifined token:" + t)

print(stack.pop())
