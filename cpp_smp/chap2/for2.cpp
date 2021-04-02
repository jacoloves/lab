#include <iostream>
using namespace std;

int score[] = {
  1, 2, 4, 4, 5, 2, 4,
  1, 2, 3, 4, 4, 5, 5,
  3, 4, 5, 5, 6, 6, 6,
};

int main() {
  int sum = 0;

  for(int i = 0; i < sizeof(score)/sizeof(score[0]); ++i) {
    sum += score[i];
  }
   
  cout << "class average is " << sum / (sizeof(score)/sizeof(score[0])) 
    << " points." << endl;
}
