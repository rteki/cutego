#include "public.h"

#include <iostream>

int init()
{
    instance = new CuteGo();
}

void start()
{
    instance->start();
}
