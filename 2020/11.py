# initial solution

I = [line.strip() for line in open('input11.txt') if line]
R = len(I)
C = len(I[0])
F = {}
for r, row in enumerate(I):
    for c, cell in enumerate(row):
        F[r, c] = cell

def neighbours1(r, c):
    ns = [
        (r-1, c),
        (r+1, c),
        (r, c+1),
        (r, c-1),
        (r+1, c+1),
        (r-1, c-1),
        (r+1, c-1),
        (r-1, c+1),
    ]
    return [(row, col) for row, col in ns if 0<=row<R and 0<=col<C]

def neighbours2(r, c):
    ns = []

    # -.
    for i in range(1, 100):
        if r-i < 0: break
        if F[r-i, c] != '.':
            ns.append((r-i, c))
            break
    # +.
    for i in range(1, 100):
        if r+i >= R: break
        if F[r+i, c] != '.':
            ns.append((r+i, c))
            break

    # .-
    for i in range(1, 100):
        if c-i < 0: break
        if F[r, c-i] != '.':
            ns.append((r, c-i))
            break
    # .+
    for i in range(1, 100):
        if c+i >= C: break
        if F[r, c+i] != '.':
            ns.append((r, c+i))
            break

    # ++
    for i in range(1, 100):
        if r+i >= R or c+i >= C: break
        if F[r+i, c+i] != '.':
            ns.append((r+i, c+i))
            break
    # --
    for i in range(1, 100):
        if r-i < 0 or c-i < 0: break
        if F[r-i, c-i] != '.':
            ns.append((r-i, c-i))
            break

    # +-
    for i in range(1, 100):
        if r+i >= R or c-i < 0: break
        if F[r+i, c-i] != '.':
            ns.append((r+i, c-i))
            break
    # -+
    for i in range(1, 100):
        if r-i < 0 or c+i >= C: break
        if F[r-i, c+i] != '.':
            ns.append((r-i, c+i))
            break

    return ns

def next_floor(F, neighbours, N):
    F2 = {}
    for r, c in F:
        occupied = len([1 for nr, nc in neighbours[r,c] if F[nr,nc] == '#'])
        if F[r,c] == 'L' and occupied == 0:
            F2[r,c] = '#'
        elif F[r,c] == '#' and occupied >= N:
            F2[r,c] = 'L'
        else:
            F2[r,c] = F[r,c]
    return F2

def solve(F, N, neighbours):
    while True:
        F2 = next_floor(F, neighbours, N)
        if F == F2:
            break
        F = F2
    return len([1 for r, c in F if F[r,c] == '#'])

# part 1
print(solve(F, 4, {(r,c):neighbours1(r, c) for r in range(R) for c in range(C)}))

# part 2
print(solve(F, 5, {(r,c):neighbours2(r, c) for r in range(R) for c in range(C)}))

