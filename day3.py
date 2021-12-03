data = '''00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010'''.split('\n')

data = open('day3.txt')
inp = [line.rstrip() for line in data]

def most_common_bit(inp, i):
    ones = sum(1 for n in inp if n[i] == '1')
    zeros = len(inp) - ones
    return '1' if ones >= zeros else '0'

# part 1
g = ''.join(most_common_bit(inp, i) for i in range(len(inp[0])))
e = ''.join('1' if b == '0' else '0' for b in g)
print(int(g, 2) * int(e, 2))

# part 2
def o(inp, i, num): return num[i] == most_common_bit(inp, i)
def c(inp, i, num): return num[i] != most_common_bit(inp, i)
def filter(inp, f):
    i = 0
    while len(inp) > 1:
        inp = [n for n in inp if f(inp, i, n)]
        i += 1
    return int(inp[0], 2)
print(filter(inp, o) * filter(inp, c))
