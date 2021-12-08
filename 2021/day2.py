inp = [(line.split()[0], int(line.split()[1])) for line in open('day2.txt')]

# part 1
d, p = 0, 0
for cmd, n in inp:
    if cmd == 'down':
        d += n
    elif cmd == 'up':
        d -= n
    elif cmd == 'forward':
        p += n

print(d*p)

# part 2
d, p, a = 0, 0, 0
for cmd, n in inp:
    if cmd == 'down':
        a += n
    elif cmd == 'up':
        a -= n
    elif cmd == 'forward':
        p += n
        d += a * n

print(d*p)