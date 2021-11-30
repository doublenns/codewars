def digits(n):
    count = 1
    while n >= 10:
        n /= 10
        count += 1
    return count