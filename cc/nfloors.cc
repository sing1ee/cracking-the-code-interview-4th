#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
	int n;
	while (cin >> n) {
		unsigned long *a = new unsigned long[55];
		a[1] = 1;
		unsigned long sum = a[1];
		for (int i = 2; i <= n; ++i) {
			a[i] = sum + 1;
			sum += a[i];
		}
		cout << a[n] << endl;
	}
	return 0;
}