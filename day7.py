from collections import Counter
from pprint import pprint

# inp = '16,1,2,0,4,2,7,1,2,14'
inp = open('day7.txt').read()
d = list(map(int, inp.split(',')))

def solve(f):
    return min(sum(f(c, p) for c in d) for p in range(min(d), max(d)+1))

def cost1(c, p): return abs(c - p)
def cost2(c, p): return sum(i for i in range(1, abs(c - p) + 1))

pprint(solve(cost1))
pprint(solve(cost2))
