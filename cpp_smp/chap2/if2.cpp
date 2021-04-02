#include <iostream>
using namespace std;

void Divide() {
  int a, b;

  cout << "input 1 > " << flush;
  cin >> a;
  
  cout << "input 2 > " << flush;
  cin >> b;

  if(b==0) {
    cout << "0 isn't divided" << endl;
  } else {
    cout << a << " / " << b << " = "
      << a / b << " ... " << a % b << endl;
  }
}
int main() {
  Divide();
  Divide();
}
