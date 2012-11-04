#include <stdio.h>

int arr[1010][1010];

int main(int argc, char const *argv[]) {
	int m, n;
	int t;
	while (scanf("%d %d %d", &m, &n, &t) != EOF) {

		for (int i = 0; i < m; ++i) {
			for (int j = 0; j < n; ++j) {
				scanf("%d", &arr[i][j]);
			}
		}
		int row = 0;
		int col = n - 1;
		bool found = false;
		while (col >= 0 && row < m) {

			if (t == arr[row][col]) {
				found = true;
				break;
			}

			if (t < arr[row][col]) {
				--col;
			} else {
				++row;
			}
		}
		if (found) {
			printf("YES\n");
		} else {
			printf("NO\n");
		}
	}
	return 0;
}