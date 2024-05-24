import datetime

start = datetime.datetime.now()

for number in range(1, 1000001):
    if number % 3 == 0:
        print(f"{number}. Fizz")
    elif number % 5 == 0:
        print(f"{number}. Buzz")
    elif number % 3 == 0 and number % 5 == 0:
        print(f"{number}. FizzBuzz")
    else:
        print(number)

timeElapsed = (datetime.datetime.now() - start).total_seconds()

print(f"Total time for execution: {timeElapsed}")
