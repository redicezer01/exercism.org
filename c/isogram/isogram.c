#include "isogram.h"
#include <stddef.h>

#define LEN 256
#define ALLOWED(c) (c == ' ' || c == '-')
#define IS_LETTER(c) (('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z'))

char to_lower(char c);

bool is_isogram(const char phrase[])
{
	int abc[LEN];
	int i, c;
	int iso;

	if (phrase == NULL)
		return false;

	iso = true;	

	for (i = 0; i < LEN; ++i)
		abc[i] = 0;	

	for (i = 0; phrase[i] != '\0'; ++i) {
		c = to_lower(phrase[i]);	
		++abc[c];
		if (abc[c]>1 && ALLOWED(c)==false)
			iso = false;
	}
		
	return iso;
}

char to_lower(char c)
{
	int r;
	r = 'a' - 'A';
	if (IS_LETTER(c) == true) {
		if (c < 'a')
			return c + r;
		else if (c > 'z')
			return c - r;
	}
	return c;
}
