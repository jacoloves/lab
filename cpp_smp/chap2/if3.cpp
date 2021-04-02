#include <iostream>
using namespace std;

int WesternToShouwa(int western) {
  if(1926 <= western && western <=1989) {
    return western - 1925;
  } else {
    return 0;
  }
}

void Shouwa() {
  int western;

  cout << "input > " << flush;
  cin >> western;

  int shouwa = WesternToShouwa(western);
  if(shouwa == 0) {
    cout << "this year is not shouwa." << endl;
  } else {
    cout << "this year is " << shouwa << " year." << endl;
  }
}
int main() {
  Shouwa();
  Shouwa();
}
