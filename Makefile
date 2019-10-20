#temp makefile

ifeq ($(OS),Windows_NT)
	RM=del
	EXT_EXE=.exe
	EXT_DL=.dll
else 
	RM=rm
	EXT_DL=.so
endif

clean:
	$(RM) main$(EXE)
	$(RM) CuteGo$(EXT_DL)