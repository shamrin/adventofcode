inp = [[int(c) for c in line.rstrip()] for line in open('day3.txt')]

def most_common(inp, i):
    return int(sum(1 for n in inp if n[i])
            >= sum(1 for n in inp if not n[i]))
def least_common(inp, i):
    return int(not most_common(inp, i))
def number(bits):
    return int(''.join(str(bit) for bit in bits), 2)

# part 1
N = len(inp[0])
g = number(most_common(inp, i) for i in range(N))
e = number(least_common(inp, i) for i in range(N))
print(g * e)

# part 2
def distill(inp, bit):
    i = 0
    while len(inp) > 1:
        inp = [n for n in inp if n[i] == bit(inp, i)]
        i += 1
    return number(inp[0])
print(distill(inp, most_common) * distill(inp, least_common))
