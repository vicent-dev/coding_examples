
def sum(n1 :float, n2 :float):
    return n1 + n2

def sub(n1 :float, n2 :float):
    return n1 - n2

def mult(n1 :float, n2 :float):
    return n1 * n2

if __name__ == "__main__":
    number1: float = 2.5 
    number2: float = 3.89

    print("{} + {} = {}".format(number1, number2, sum(number1, number2)))
    print("{} - {} = {}".format(number1, number2, sub(number1, number2)))
    print("{} * {} = {}".format(number1, number2, mult(number1, number2)))
