# inspired by https://www.reddit.com/r/adventofcode/comments/kaw6oz/2020_day_11_solutions/gfd91ts/

I = ['X' + line.strip() + 'X' for line in open('input11.txt') if line]

I = ['X' * len(I[0])] + I + ['X' * len(I[0])]
R = len(I)
C = len(I[0])

F = {}
for r, row in enumerate(I):
    for c, cell in enumerate(row):
        F[r, c] = cell

def solve(F, ray, N):
    while True:
        F2 = F.copy()
        for r in range(1,R-1):
            for c in range(1,C-1):
                occupied = len([1 for dr in (-1,0,1) for dc in (-1,0,1)
                    if (dr, dc) != (0,0) and ray(F, r, c, dr, dc) == '#'])
                if F[r,c] == 'L' and occupied == 0:
                    F2[r,c] = '#'
                elif F[r,c] == '#' and occupied >= N:
                    F2[r,c] = 'L'
                else:
                    F2[r,c] = F[r,c]
        if F2 == F:
            break
        F = F2
    return len([1 for r, c in F if F[r,c] == '#'])

# part 1
print(solve(F, lambda F,r,c,dr,dc: F[r+dr,c+dc], 4))

# part 2
def ray(F, r, c, dr, dc):
    r += dr
    c += dc
    while F[r, c] == '.':
        r += dr
        c += dc
    return F[r, c]
print(solve(F, ray, 5))
