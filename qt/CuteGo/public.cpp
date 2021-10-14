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

void registerResource(char * path) {
    instance->registerResource(QString::fromStdString( std::string(path) ));
}

void loadQmlEntry(char *path)
{
    instance->loadQmlEntry( QString::fromStdString( std::string(path) ) );
}

void newEventManager(char *name)
{
    instance->newEventManager( QString::fromStdString( std::string(name) ) );
}

void callQt(char * eventManagerName, char * eventName, char * strValue)
{
    instance->call(QString::fromStdString( std::string(eventManagerName) ), QString::fromStdString( std::string(eventName) ), QString::fromStdString( std::string(strValue) ));
}
