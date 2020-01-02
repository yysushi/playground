#include <iostream>
#include <iomanip> // for setprecision

int main() {
  std::cout << 123.456 <<  std::endl;
  std::cout << std::fixed << std::setprecision(2) << 123.456 << std::endl;
  return 0;
}
