#include <iostream>
using namespace std;

int f(int x, int y);

int main() {
  cout << "f(1, 2) = " << f(1,2) << endl;
  cout << "f(181, 222) = " << f(181,222) << endl;
}

int f(int x, int y) {
  return 2 * x + y;
}
