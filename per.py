MAX_CHAR = 26

def factorial(n) :
    fact = 1;
    for i in range(2, n + 1) :
        fact = fact * i;
    return fact

def countDistinctPermutations(st) :  
    length = len(st)
    freq = [0] * MAX_CHAR
  
    for i in range(0, length) :
        if (st[i] >= 'A') :
            freq[(ord)(st[i]) - 65] = freq[(ord)(st[i]) - 65] + 1;
  
    fact = 1
    for i in range(0, MAX_CHAR) :
        fact = fact * factorial(freq[i])

    return int(factorial(length) / fact)

def main():
  word = input()
  print(countDistinctPermutations(word))

main()