#include "grains.h"

uint64_t square(uint8_t index)
{
	
	if (index <= 0 || index > 64)
		return 0;

	uint64_t sq = 1;
	return (sq << (index-1));

}

uint64_t total(void)
{
	uint64_t t = 0;
	return ~(t);
}
