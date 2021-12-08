inputs = [
'0,3,6',
'16,11,15,0,1,7',
]

def solve(ns, steps):
    last, c = ns[-1], {n:i+1 for i, n in enumerate(ns)}
    for n in range(len(ns), steps):
        c[last], last = n, n - c.get(last, n)
    return last

for inp in inputs:
    ns = [int(n) for n in inp.split(',')]
    for steps in (2020, 30000000):
        print(f'{steps}: {solve(ns, steps)}')
