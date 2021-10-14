#ifndef CUTEGO_H
#define CUTEGO_H


#include <QtGui/QGuiApplication>
#include <QtQml/QQmlApplicationEngine>
#include <QString>
#include <QMap>

#include "EventManager.h"

typedef void (*GoCallbackFunc)(const char*, const char*, const char*);

class CuteGo
{
public:
    CuteGo(GoCallbackFunc gocallback);
    void start();
    void registerResource(QString path);
    void loadQmlEntry(QString qmlEntry);
    void newEventManager(QString name);
    void call(QString eventManagerName, QString eventName, QString value);



private:
    QGuiApplication *app;
    QQmlApplicationEngine *engine;
    QMap<QString, EventManager*> eventEmitters;
    GoCallbackFunc gc;
};

#endif // CUTEGO_H
