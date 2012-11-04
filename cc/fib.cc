#include <iostream>
#include <string.h>

using namespace std;

unsigned long *a = new unsigned long[100];

int main(int argc, char const *argv[])
{
	int n;
	while (cin >> n) {
		memset(a, 0, sizeof(a));
		a[0] = 0;
		a[1] = 1;
		for (int i = 2; i <= n; ++i) {
			a[i] = a[i - 1] + a[i - 2];
		}
		cout << a[n] << endl;
	}
	return 0;
}