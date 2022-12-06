DIR_Config   = ./IT8951-ePaper/Raspberry/lib/Config
DIR_EPD      = ./IT8951-ePaper/Raspberry/lib/e-Paper
DIR_FONTS    = ./IT8951-ePaper/Raspberry/lib/Fonts
DIR_GUI      = ./IT8951-ePaper/Raspberry/lib/GUI

DIR_BIN      = ./bin

OBJ_C = $(wildcard ${DIR_Config}/*.c ${DIR_EPD}/*.c ${DIR_FONTS}/*.c ${DIR_GUI}/*.c ${DIR_Examples}/*.c )
OBJ_O = $(patsubst %.c,${DIR_BIN}/%.o,$(notdir ${OBJ_C}))

TARGET = epd

CC = gcc

MSG = -g -O0 -Wall
DEBUG = -D USE_DEBUG
STD = -std=gnu99

CFLAGS += $(MSG) $(DEBUG) $(STD)

LIB = -lbcm2835 -lm -lrt -lpthread

$(shell mkdir -p $(DIR_BIN))

${TARGET}:${OBJ_O}
	$(CC) $(CFLAGS) $(OBJ_O) -o $@ $(LIB) 

${DIR_BIN}/%.o:$(DIR_Config)/%.c
	$(CC) $(CFLAGS) -c  $< -o $@ 

${DIR_BIN}/%.o:$(DIR_EPD)/%.c
	$(CC) $(CFLAGS) -c  $< -o $@ 
	
${DIR_BIN}/%.o:$(DIR_FONTS)/%.c
	$(CC) $(CFLAGS) -c  $< -o $@ 
	
${DIR_BIN}/%.o:$(DIR_GUI)/%.c
	$(CC) $(CFLAGS) -c  $< -o $@ 

clean :
	rm $(DIR_BIN)/*.* 
	rm $(TARGET) 

