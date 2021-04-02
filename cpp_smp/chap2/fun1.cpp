#include <iostream>
using namespace std;

int f(int x, int y){
  return 2 * x + y;
}

int main() {
  cout << "f(1, 2) = " << f(1, 2) << endl;
  cout << "f(182, 144) = " << f(182, 144) << endl;
}
