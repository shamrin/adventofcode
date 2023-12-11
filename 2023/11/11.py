from itertools import combinations

def solve(gs, k):
    nr = max(r for r, _ in gs) + 1
    nc = max(c for _, c in gs) + 1
    grows = {r for r, _ in gs}
    gcols = {c for _, c in gs}
    emptyrows = sorted(r for r in range(nr) if r not in grows)
    emptycols = sorted(c for c in range(nc) if c not in gcols)
    for er in reversed(emptyrows):
        gs = {(r+k-1 if r > er else r, c) for r, c in gs}
    for ec in reversed(emptycols):
        gs = {(r, c+k-1 if c > ec else c) for r, c in gs}
    return sum(abs(r1 - r2) + abs(c1 - c2) for (r1, c1), (r2, c2) in combinations(gs, 2))

inp = open('11.txt').read()
gs = {(r, c) for r, L in enumerate(inp.split("\n")) for c, ch in enumerate(L) if ch == '#'}
print(solve(gs, 2))
print(solve(gs, 10**6))
