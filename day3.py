inp = [line.rstrip() for line in open('day3.txt')]

def most_common(inp, i):
    return sum(1 for n in inp if n[i] == '1') \
        >= sum(1 for n in inp if n[i] == '0')
def least_common(inp, i):
    return not most_common(inp, i)
def s(bit): return str(int(bit))

# part 1
N = len(inp[0])
g = ''.join(s(most_common(inp, i)) for i in range(N))
e = ''.join(s(least_common(inp, i)) for i in range(N))
print(int(g, 2) * int(e, 2))

# part 2
def distill(inp, bit):
    i = 0
    while len(inp) > 1:
        inp = [n for n in inp if n[i] == s(bit(inp, i))]
        i += 1
    return int(inp[0], 2)
print(distill(inp, most_common) * distill(inp, least_common))
