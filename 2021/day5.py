# %%
inp='''0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2'''

# %%
from collections import Counter
from itertools import groupby

inp = open('day5.txt').read()
lines = [[[int(n) for n in xy.split(',')] for xy in line.split(' -> ')] for line in inp.split('\n')]

d = Counter()
for (x1, y1), (x2, y2) in lines:
    if x1 == x2:
        ys = 1 if y2 > y1 else -1
        d.update((x1, y) for y in range(y1, y2+ys, ys))
    elif y1 == y2:
        xs = 1 if x2 > x1 else -1
        d.update((x, y1) for x in range(x1, x2+xs, xs))
    else:
        # diagonal
        ys = 1 if y2 > y1 else -1
        xs = 1 if x2 > x1 else -1
        d.update(zip(range(x1, x2 + xs, xs), range(y1, y2 + ys, ys)))


print(sum(1 for k, c in d.items() if c >= 2))
