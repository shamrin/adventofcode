from collections import Counter
from itertools import product
I = '...#..#.\n#..#...#\n.....###\n##....##\n......##\n........\n.#......\n##...#..'
g = {(x,y,0,0) for y, row in enumerate(I.split('\n')) for x, c in enumerate(row) if c == '#'}
def add(v1, v2): return tuple(e1+e2 for e1, e2 in zip(v1,v2))
def next(g, Dw, D=(-1,0,1)):
    ns = Counter(add(p, d) for p in g for d in product(D,D,D,Dw) if d != (0,0,0,0))
    return {p for p, n in ns.items() if n == 3 or (n == 2 and p in g)}
for _ in range(6): g = next(g, (-1,0,1)) # for part 1, replace with `next(g, (0,))`
print(len(g))
