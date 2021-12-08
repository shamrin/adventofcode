from collections import Counter

inp = open('day6.txt').read()
F = list(map(int, inp.split(',')))

def gen(days):
    D = Counter(F)
    for _ in range(days):
        D2 = Counter()
        for f, n in D.items():
            if f == 0:
                D2[8] += n
                D2[6] += n
            else:
                D2[f - 1] += n
        D = D2

    return sum(D.values())

# part 1
print(gen(80))
# part 2
print(gen(256))