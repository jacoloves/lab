#include <iostream>
#include <cstdio>
using namespace std;

int f(int x, int y) {
  return 2 * x + y;
}

void show(int x, int y){
  char str[50];

  sprintf(str, "f(%d, %d) = %d", x, y, f(x, y));
  cout << str << endl;
}


int main() {
  show(1, 2);
  show(182, 144);
}
