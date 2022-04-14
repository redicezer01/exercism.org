#include "resistor_color.h"
#include <stdlib.h>


int color_code(resistor_band_t v)
{
	return (int) v;
}

/* static version */
resistor_band_t *colors()
{
	
	static resistor_band_t r[] = {BLACK, BROWN, RED, ORANGE, YELLOW, GREEN, BLUE, VIOLET, GREY, WHITE};
	return r;
}
