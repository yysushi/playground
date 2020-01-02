#include <iostream>
#include <math.h>

float round_n(float a, int n) {
  return round(pow(10, n - 1) * a) / pow(10, n - 1);
}

int main() {
  float a = 0.555;
  std::cout << round(a) << std::endl;
  std::cout << round(pow(10, 1) * a) / 10 << std::endl;
  std::cout << round_n(a, 2) << std::endl;
  std::cout << round_n(a, 3) << std::endl;
  return 0;
}
