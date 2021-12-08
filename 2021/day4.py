# %%
inp = '''7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7'''

# %%
inp = open('day4.txt').read()
draw, _, boards = inp.partition('\n\n')

nboards = len(boards.split('\n\n'))

B = {
    (b, r, c): int(n)
    for b, board in enumerate(boards.split('\n\n'))
        for r, line in enumerate(board.split('\n'))
            for c, n in enumerate(line.split())
}

# %%
def part1():
    M = {}
    def winner_board():
        for b in range(nboards):
            for r in range(5):
                if all(M.get((b, r, c)) for c in range(5)):
                    return b
            for c in range(5):
                if all(M.get((b, r, c)) for r in range(5)):
                    return b

    for d in draw.split(','):
        d = int(d)
        for k in B:
            if B[k] == d:
                M[k] = True
        if b := winner_board():
            s = sum(B[b, r, c] for r in range(5) for c in range(5) if not M.get((b, r, c)))
            print(s * d)
            break

part1()

# %%
def part2():
    M = {}
    def winner_boards(bs):
        w = set()
        for b in bs:
            for r in range(5):
                if all(M.get((b, r, c)) for c in range(5)):
                    w.add(b)
            for c in range(5):
                if all(M.get((b, r, c)) for r in range(5)):
                    w.add(b)
        return w

    bs = set(range(nboards))
    for d in draw.split(','):
        d = int(d)
        for k in B:
            if B[k] == d:
                M[k] = True
        if (b := winner_boards(bs)) != set():
            if len(bs) == 1:
                (b,) = b
                s = sum(B[b, r, c] for r in range(5) for c in range(5) if not M.get((b, r, c)))
                print(s * d)
                break
            bs -= b

part2()