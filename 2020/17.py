from collections import Counter
from itertools import product

I='''
...#..#.
#..#...#
.....###
##....##
......##
........
.#......
##...#..
'''.strip()

D = (-1,0,1)
def add(v1, v2): return tuple(e1+e2 for e1, e2 in zip(v1,v2))
def solve(I, dim):
    g = {(x,y) + (0,)*(dim-2)
         for y, r in enumerate(I.split('\n')) for x, c in enumerate(r) if c == '#'}
    for _ in range(6):
        ns = Counter(add(p, d) for p in g for d in product(*[D]*dim) if any(d))
        g = {p for p, n in ns.items() if n == 3 or (n == 2 and p in g)}
    return len(g)

print(solve(I, 2))
print(solve(I, 3))
print(solve(I, 4))
print(solve(I, 5))
