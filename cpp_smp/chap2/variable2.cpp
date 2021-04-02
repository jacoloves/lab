#include <iostream>
using namespace std;

int main() {
  int a = 0;

  a += 2;
  cout << "a += 2          -> a = " << a << endl;

  ++a;
  cout << "++a             -> a = " << a << endl;

  a = a * 5 + 2;
  cout << "a = a * 5 + 2   -> a = " << a << endl;

}
