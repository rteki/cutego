#ifndef PUBLIC_H
#define PUBLIC_H

#ifdef WIN32
#define EXPORT __declspec(dllexport)
#define IMPORT __declspec(dllimport)
#endif

#ifdef __linux__
#define EXPORT __attribute__((visibility("default")))
#define IMPORT
#endif



#ifdef CUTEGO_LIBRARY

#include "CuteGo.h"

#define EXTERN_C extern "C"
#define DL EXPORT

#else

#define EXTERN_C
#define DL IMPORT

#endif

typedef void (*GoCallbackFunc)(const char*, const char*, const char*);

//Initializing CuteGo object
EXTERN_C DL void init(GoCallbackFunc gocallback);


#ifdef CUTEGO_LIBRARY
static CuteGo* instance;
#endif

EXTERN_C DL void start();
EXTERN_C DL void loadQmlEntry(char * path);
EXTERN_C DL void newEventManager(char * name);
EXTERN_C DL void callQt(char * eventManagerName, char * eventName, char * strValue);

#endif // PUBLIC_H
