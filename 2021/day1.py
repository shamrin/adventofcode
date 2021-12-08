N = [int(i) for i in open('day1.txt').read().split()]

# part 1
increased = 0
for i in range(1, len(N)):
    if N[i] > N[i-1]:
        increased += 1
print(increased)

# part 2
increased = 0
def window(i):
    return N[i] + N[i+1] + N[i+2]
for i in range(1, len(N)-2):
    if window(i) > window(i-1):
        increased += 1
print(increased)