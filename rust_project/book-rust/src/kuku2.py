for y in range(1, 10):
    a = ["{:3}".format(x * y) for x in range(1, 10)]
    print(",".join(a))
