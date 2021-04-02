#include <iostream>
using namespace std;

void Showacode(char ch) {
  cout << "[" << ch << "] is "
    << (int)(unsigned char)ch << " ." << endl;
}
int main() {
  Showacode('0');
  Showacode('A');
}
