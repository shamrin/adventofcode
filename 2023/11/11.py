input = open('11.txt').read()

gals = set()
i = 1
for r, line in enumerate(input.split("\n")):
    for c, cell in enumerate(line):
        if cell == "#":
            gals.add((i, r, c))
            i+=1

nr = len(input.split("\n"))
nc = len(input.split("\n")[0])

galrs = set(r for i, r, c in gals)
galcs = set(c for i, r, c in gals)

ecs = sorted(c for c in range(nc) if c not in galcs)
ers = sorted(r for r in range(nr) if r not in galrs)

k = 10**6

for ec in reversed(ecs):
    gals = set((i, r, ((c+k-1) if c > ec else c)) for i, r, c in gals)
for er in reversed(ers):
    gals = set((i, ((r+k-1) if r > er else r), c) for i, r, c in gals)

from itertools import combinations

s = 0
for (i1, r1, c1), (i2, r2, c2) in combinations(gals, 2):
    d = abs(r1 - r2) + abs(c1 - c2)
    s += d
print(s)
