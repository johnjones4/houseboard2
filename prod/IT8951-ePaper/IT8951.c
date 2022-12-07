#include "IT8951.h"
#include "EPD_IT8951.h"

#define VCOM 16
IT8951_Dev_Info devInfo = {0, 0};
UDOUBLE Init_Target_Memory_Addr;

void ext_IT8951_init() {
    DEV_Module_Init();
    devInfo = EPD_IT8951_Init(VCOM);
    Init_Target_Memory_Addr = devInfo.Memory_Addr_L | (devInfo.Memory_Addr_H << 16);
    EPD_IT8951_Clear_Refresh(devInfo, Init_Target_Memory_Addr, INIT_Mode);
}

UWORD ext_IT8951_width() {
    return devInfo.Panel_W;
}

UWORD ext_IT8951_height() {
    return devInfo.Panel_W;
}

void ext_IT8951_draw(UBYTE* image) {
    EPD_IT8951_4bp_Refresh(image, 0, 0, devInfo.Panel_W, devInfo.Panel_W, false, Init_Target_Memory_Addr,false);
}
