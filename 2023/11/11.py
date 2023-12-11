from itertools import combinations

inp = open('11.txt').read()
k = 10**6

gs = {(r, c) for r, line in enumerate(inp.split("\n")) for c, cell in enumerate(line) if cell == '#'}

nr = len(inp.split("\n"))
nc = len(inp.split("\n")[0])

grows = set(r for r, _ in gs)
gcols = set(c for _, c in gs)

emptyrows = sorted(r for r in range(nr) if r not in grows)
emptycols = sorted(c for c in range(nc) if c not in gcols)

for er in reversed(emptyrows):
    gs = set((r+k-1 if r > er else r, c) for r, c in gs)
for ec in reversed(emptycols):
    gs = set((r, c+k-1 if c > ec else c) for r, c in gs)

print(sum(abs(r1 - r2) + abs(c1 - c2) for (r1, c1), (r2, c2) in combinations(gs, 2)))
