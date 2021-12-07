inp = open('day7.txt').read()
d = list(map(int, inp.split(',')))

def solve(cost): return min(sum(cost(abs(c - p)) for c in d) for p in range(min(d), max(d)+1))
def cost1(dist): return dist
def cost2(dist): return (1 + dist) * dist // 2

print(solve(cost1))
print(solve(cost2))
