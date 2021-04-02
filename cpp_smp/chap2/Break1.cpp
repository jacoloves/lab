#include <iostream>
using namespace std;

int main() {
  for(;;) {
    int a, b;

    cout << "1 input > " << flush;
    cin >> a;

    cout << "2 input > " << flush;
    cin >> b;

    if(b == 0) {
      cout << "0 didn't divided. " << endl;
      break;
    }

    cout << a << " / " << b << " = "
      << a / b << "..." << a % b << endl;
  }

  cout << "program is finished." << endl;
}
