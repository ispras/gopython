import random

class unique_python_calc:
    def __init__(self):
        print("__init__() called")
        self.unique_number = random.randint(-10, 10)
    
    def calc(self, number):
        print("calc() called with arg:", number)
        
        # unique calc with unique python tool...
        res = number * self.unique_number

        print("calc() res =", res)
        return res
