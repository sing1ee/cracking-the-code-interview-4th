#include <stdio.h>

int main(int argc, char const *argv[])
{
	int n, k;
	while (scanf("%d %d", &n, &k) != EOF) 
	{
		int max = 0;
		int *a = new int[n];
		for (int i = 0; i < n; ++i) {
			scanf("%d", &a[i]);
			if (a[i] > max) {
				max = a[i];
			}
		}

		int first = 0;
		int last = n - 1;
		int mul = max * max;
		int ff = -1;
		int fl = -1;
		while (first < last) {
			int sum = a[first] + a[last];
			if (sum == k) {
				if (a[first] * a[last] < mul) {
					mul = a[first] * a[last];
					ff = first;
					fl = last;
					first++;
					last--;
				}
			} else if (sum < k) {
				first++;
			} else {
				last--;
			}
		}
		if (-1 != ff) {
			printf("%d %d\n", a[ff], a[fl]);
		} else {
			printf("-1 -1\n");
		}
	}
	return 0;
}