#include <iostream>
using namespace std;

int main() {
  for(int i=0; i<5; ++i) {
    int a, b, op, result;

    cout << "1 > " << flush;
    cin >> a;
    cout << "2 > " << flush;
    cin >> b;
    cout << "1:plus 2:minus 3:multi 4:div > " << flush;
    cin >> op;

    switch(op) {
      case 1:
        result = a+b; break;
      case 2:
        result = a-b; break;
      case 3:
          result = a*b; break;
      case 4:
          result = a/b; break;
      default:
          cout << "plese input 1, 2, 3, 4" << endl;
          continue;
    }

    cout << "answer is " << result << " . " << endl;
  }
}
