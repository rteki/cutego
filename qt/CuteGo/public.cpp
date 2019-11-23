#include "public.h"

#include <iostream>

void init(GoCallbackFunc gocallback)
{
    instance = new CuteGo(gocallback);
}

void start()
{
    instance->start();
}

void loadQmlEntry(char *path)
{
    instance->loadQmlEntry( QString::fromStdString( std::string(path) ) );
}

void newEventManager(char *name)
{
    instance->newEventManager( QString::fromStdString( std::string(name) ) );
}
