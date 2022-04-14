#include "hamming.h"

int lens(const char *s);

int compute(const char *lhs, const char *rhs)
{
	int hd = -1; /* Hamming Distance
			0: equals or both empty
		       >0 - < length of strands:
				number of diffrences
		       -1: diffrent length of lhs and rhs */

	if (lens(lhs) != lens(rhs))
		return hd;
	hd = 0;

	while (*lhs != '\0') {
		if (*lhs++ != *rhs++)
			hd++;
	}

	return hd;
}

/* gets length (count of characters) of s */
int lens(const char *s)
{
	int i;

	for (i = 0; *s++ != '\0'; ++i)
		;

	return i;
}
