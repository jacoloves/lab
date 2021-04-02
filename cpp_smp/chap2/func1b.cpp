#include <iostream>
using namespace std;

int f(int x, int y){
  return 2 * x + y;
}

void show(int x, int y) {
  cout << "f(" << x << ", " << y << ") = " << f(x, y) << endl;
}

int main() {
  show(1, 2);
  show(123, 999);
}
