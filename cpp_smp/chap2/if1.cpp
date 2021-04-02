#include <iostream>
using namespace std;

void BirthMonth() {
  int month;

  cout << "What day is your birthday? > " << flush;
  cin >> month;

  if (1 <= month && month <= 12) {
  cout << "oh, " << month << " ." << endl;
  } else {
    cout << month << " what?" << endl
         << "This month is none." << endl;
  }
}
int main() {
  BirthMonth();
  BirthMonth();
}
