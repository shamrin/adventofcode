inp = open('13.txt').read()

def rowsym(nr, nc, g, row):
    for c in range(0, nc):
        for r in range(0, min(row, nr - row)):
            if ((row + r, c) in g) != ((row-r-1, c) in g):
                return False
    return True

def summarize(grid, *, smudge = False):
    lines = grid.split("\n")
    nr = len(lines)
    nc = len(lines[0])
    g, gt = set(), set()
    for r, line in enumerate(lines):
        for c, ch in enumerate(line):
            if ch == '#':
                g.add((r, c))
                gt.add((c, r))

    result = 0
    for r in range(1, nr):
        if rowsym(nr, nc, g, r):
            result = 100 * r
    for c in range(1, nc):
        if rowsym(nc, nr, gt, c):
            result = c

    if not smudge:
        return result

    for rs in range(0, nr):
        for cs in range(0, nc):
            # fix smudge
            removed = False
            if (rs, cs) in g:
                removed = True
                g.remove((rs, cs))
                gt.remove((cs, rs))
            else:
                g.add((rs, cs))
                gt.add((cs, rs))

            for r in range(1, nr):
                if result == r * 100: continue
                if rowsym(nr, nc, g, r):
                    return 100 * r
            for c in range(1, nc):
                if result == c: continue
                if rowsym(nc, nr, gt, c):
                    return c

            # de-fix smudge
            if removed:
                g.add((rs, cs))
                gt.add((cs, rs))
            else:
                g.remove((rs, cs))
                gt.remove((cs, rs))

    return result

print(sum(summarize(grid, smudge=False) for grid in inp.split("\n\n")))
print(sum(summarize(grid, smudge=True) for grid in inp.split("\n\n")))
