inp = open('input16.txt').read().strip()

specs, your, tickets = inp.split('\n\n')
your = your.partition('\n')[2]
tickets = tickets.partition('\n')[2]

tickets = [[int(v) for v in ticket.split(',')] for ticket in tickets.split('\n')]
specs = [(s.partition(': ')[0], s.partition(': ')[2].split(' or ')) for s in specs.split('\n')]
specs = [(n, [(int(r.split('-')[0]), int(r.split('-')[1])) for r in rs]) for n, rs in specs]
your = [int(n) for n in your.split(',')]

all_ranges = []
for n, ranges in specs:
    all_ranges.extend(ranges)

def valid(v, ranges):
    return any(start <= v <= end for start, end in ranges)

# part 1
print(sum(v for ticket in tickets for v in ticket if not valid(v, all_ranges)))

# part 2
tickets = [your] + [ticket for ticket in tickets if all(valid(v, all_ranges) for v in ticket)]
candidates = [set(n for n, _ in specs) for _ in range(len(specs))]

for ticket in tickets:
    for i, v in enumerate(ticket):
        candidates[i] &= set(name for name, ranges in specs if valid(v, ranges))

while any(len(cs) > 1 for cs in candidates):
    for c in (next(iter(cs)) for cs in candidates if len(cs) == 1):
        for cs in candidates:
            if len(cs) > 1:
                cs.discard(c)
candidates = [c.pop() for c in candidates]

from functools import reduce
print(reduce(lambda a,b: a*b, (your[i] for i, c in enumerate(candidates) if c.startswith('departure'))))
