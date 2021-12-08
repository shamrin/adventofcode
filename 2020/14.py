import re

inp = [l.strip() for l in open('input14.txt') if l]
I = []
for line in inp:
    if line.startswith('mask = '):
        I.append(('mask', line[7:]))
    else:
        m = re.match(r'mem\[(\d+)\] = (\d+)', line)
        assert m
        I.append(('mem', (int(m[1]), int(m[2]))))

# part 1
mask_or = 0
mask = ''
mask_and = int('1'*36,2)
mem = {}
for cmd, data in I:
    if cmd == 'mask':
        mask = data
    elif cmd == 'mem':
        addr, r = data
        r &= int(mask.replace('X', '1'), 2) # set 0s
        r |= int(mask.replace('X', '0'), 2) # set 1s
        mem[addr] = r
print(sum(m for m in mem.values()))


def bits(n, l):
    return list(reversed(bin(n)[2:].zfill(l)))

# part 2, initial solution
mask = None
mem = {}
for cmd, data in I:
    if cmd == 'mask':
        mask = ''.join(reversed(data))
    elif cmd == 'mem':
        addr, r = data
        assert mask is not None
        for count in range(2 ** mask.count('X')):
            a = bits(addr, 36)
            c = bits(count, mask.count('X'))
            for i, m in enumerate(mask):
                if m == '1':
                    a[i] = '1'
                elif m == 'X':
                    a[i] = c[mask.count('X', 0, i)]
            mem[int(''.join(reversed(a)), 2)] = r
print(sum(m for m in mem.values()))

# part 2, alternative solution
from itertools import combinations
mask = None
mem = {}
for cmd, data in I:
    if cmd == 'mask':
        mask = ''.join(reversed(data))
    elif cmd == 'mem':
        addr, r = data
        assert mask is not None
        x_indices = [i for i, b in enumerate(mask) if b == 'X']
        for n in range(mask.count('X') + 1):
            for indices in combinations(x_indices, n):
                a = [
                    b if m == '0' else '1' if m == '1' or i in indices else '0'
                    for i, (b, m) in enumerate(zip(bits(addr, 36), mask))
                ]
                mem[int(''.join(reversed(a)), 2)] = r
print(sum(m for m in mem.values()))
