#include "grains.h"
#include <limits.h>
uint64_t square(uint8_t index)
{
    return (index > 0 && index < 65) ? 1ul << (index - 1) : 0;
}
uint64_t total(void)
{
    return UINT64_MAX;
}
