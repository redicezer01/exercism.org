#include <stdio.h>
#include <time.h>
#include "grains.h"

int grains = 1;

#ifndef GRAINS_H
grains = 0;
#endif

int main()
{
	clock_t st, et;
	long long tt = 0;
	long int i;

	for (i = 0; i < 1e+6; ++i) {
		st = clock();
		
		square(64);

		et = clock();
		tt += et - st;
	}

	printf("time: %.9f\n", (double) tt / CLOCKS_PER_SEC / 1e+6);

	return (0);
}
