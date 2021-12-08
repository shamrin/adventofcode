T, B = open('input13.txt').read().strip().split('\n')
T = int(T)
B = [(i, int(b)) for i, b in enumerate(B.split(',')) if b != 'x']

# part 1
ts = []
for _, b in B:
    d, r = divmod(T, b)
    assert r > 0, 'oops'
    ts.append((b, (d + 1) * b))
b, t = min(ts, key=lambda t: t[1])
print(b * (t - T))

# part 2
from functools import reduce

def prod(ns):
    return reduce(lambda a,b: a*b, ns, 1)

def solve(B, t, step):
    while not all((t + i) % b == 0 for i, b in B):
        t += step
    return t

t = 1
for i in range(1, len(B)+1):
    t = solve(B[:i], t, prod(b for _, b in B[:i-1]))
print(t)
