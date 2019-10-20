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




EXTERN_C DL void init();




#endif // PUBLIC_H
