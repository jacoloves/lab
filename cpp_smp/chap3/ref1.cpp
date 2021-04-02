#include <iostream>
using namespace std;


void WesternShouwa(int& x) {
  if(1926 <= x && x <= 1989) {
    x -= 1925;
  } else {
    x = 0;
  }
}

void SHouwa() {
  int year;

  cout << "input > " << flush;
  cin >> year;

  WesternShouwa(year);
  if(year != 0) {
    cout << "this year is " << year
      << "." << endl;
  } else {
    cout << "not shouwa" << endl;
  }
}

int main() {
  SHouwa();
  SHouwa();
}
