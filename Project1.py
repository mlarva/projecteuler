#If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
#Find the sum of all the multiples of 3 or 5 below 1000.

def multipleFinder(min, max, *multiples):
    multiplesSet = set()
    if min >= max:
        print("Minimum is not smaller than maximum")
        return
    if min < 0 or max < 1:
        print("Cannot enter negative numbers")
        return
    for x in range(min,max):
        for y in multiples:
            if x % y == 0:
                multiplesSet.add(x)
    return multiplesSet

def sumSet(setToSum):
    sum = 0
    for x in setToSum:
            sum += x
    return sum

multiplesSet = multipleFinder(1,1000, 3, 5)
print(sumSet(multiplesSet))
