void main() {
  Stopwatch stopWatch = new Stopwatch()..start();

  for (int number = 1; number <= 1000000; number++) {
    if (number % 3 == 0) {
      print("$number. Fizz");
    } else if (number % 5 == 0) {
      print("$number. Buzz");
    } else if (number % 3 == 0 && number % 5 == 0) {
      print("$number. FizzBuzz");
    } else {
      print("$number");
    }
  }

  Duration timeElapsed = stopWatch.elapsed;

  print('Total time for execution: ${timeElapsed}');
}
