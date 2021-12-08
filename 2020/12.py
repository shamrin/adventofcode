I = [(l[0], int(l[1:])) for l in open('input12.txt') if l]

def rotate(d, x, y):
    if d == 'R': return y, -x
    if d == 'L': return -y, x

def move(d, n, x, y):
    if d == 'N': y += n
    elif d == 'S': y -= n
    elif d == 'E': x += n
    elif d == 'W': x -= n
    return x, y

# part 1
xs = ys = 0
xw, yw = 1, 0
for c, n in I:
    if c in 'NSEW':
        xs, ys = move(c, n, xs, ys)
    elif c == 'F':
        xs += n * xw
        ys += n * yw
    elif c in 'LR':
        assert n % 90 == 0
        for i in range(int(n / 90)):
            xw, yw = rotate(c, xw, yw)
print(abs(xs) + abs(ys))

# part 2
xs = ys = 0
xw, yw = 10, 1
for c, n in I:
    if c in 'NSEW':
        xw, yw = move(c, n, xw, yw)
    elif c == 'F':
        xs += n * xw
        ys += n * yw
    elif c in 'LR':
        assert n % 90 == 0
        for i in range(int(n / 90)):
            xw, yw = rotate(c, xw, yw)
print(abs(xs) + abs(ys))
