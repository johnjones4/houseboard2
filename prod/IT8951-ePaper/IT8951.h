#ifndef IT8951_H
#define IT8951_H

#include "DEV_Config.h"

void ext_IT8951_init();
UWORD ext_IT8951_width();
UWORD ext_IT8951_height();
void ext_IT8951_draw(UBYTE* image);

#endif
