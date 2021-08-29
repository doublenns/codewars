def reverse_words1(s):
    words = s.split()
    i = 0; j = len(words) -1
    while i<j:
        words[i], words[j] = words[j], words[i]
        i+=1; j-=1
    return " ".join(words)

def reverse_words2(s):
    words = s.split()
    reversedWords = words[::-1]
    return " ".join(reversedWords)

def reverse_words3(s):
    words = s.split()
    return " ".join(reversed(words))