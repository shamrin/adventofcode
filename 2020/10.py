I = [int(line) for line in  open('input10.txt') if line]

I.sort()
I = [0] + I + [max(I) + 3]

# part 1
ones = 0
threes = 0
for i in range(1, len(I)):
    d = I[i] - I[i-1]
    if d == 1:
        ones += 1
    if d == 3:
        threes += 1
print(ones*threes)

# part 2
ds = [0] + [I[i] - I[i-1] for i in range(1, len(I))]
c = 1
ones = 0
for d in ds:
    if d == 1:
        ones += 1
    else:
        if ones == 2:
            c *= 2
        elif ones == 3:
            c *= 4
        elif ones == 4:
            c *= 7
        elif ones > 4:
            print('oops', ones)
        ones = 0
print(c)

# part 2, memoization solution
# inspired by https://github.com/vincentvanderweele/adventofcode-2020/blob/main/day10.js
cache = {}
def count(prev, I):
    if len(I) == 1 and I[0] - prev <= 3:
        return 1
    if I[0] - prev > 3:
        return 0
    if (prev, len(I)) not in cache:
        cache[prev, len(I)] = count(prev, I[1:]) + count(I[0], I[1:])
    return cache[prev, len(I)]
print(count(I[0], I[1:]))
