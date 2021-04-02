#include <iostream>
using namespace std;

int main() {
  int a;

  cout << "input parameter > " << flush;
  cin >> a;

  cout << "this number divided by 3 is" << a % 3 << "." << endl;
}
