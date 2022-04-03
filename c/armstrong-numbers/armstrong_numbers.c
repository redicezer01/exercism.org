#include "armstrong_numbers.h"

int sum_of_digits_power(int n);
int count_of_digits(int n);

bool is_armstrong_number(int num)
{
	if (sum_of_digits_power(num) == num)
		return true;
	else 
		return false;
}

int sum_of_digits_power(int n)
{
	int nd;		/* count of digits */
	int m;		/* number reduced after last digit taken */
	int last;	/* last digit of the number */
	int sum;

	nd = count_of_digits(n);
	sum = 0;	
	while ( n > 0) {
		m = n / 10;
		last = 	n - (m*10);
		sum = sum + pow(last, nd);
		n = m;	/* reducing number to prevent infinite cycle */
	}	
	return sum;
}

int count_of_digits(int n)
{
	int i;
	for (i = 0; n > 0; ++i)
		n = n / 10;
	return i;
}

