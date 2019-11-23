#ifndef CUTEGO_H
#define CUTEGO_H


#include <QtGui/QGuiApplication>
#include <QtQml/QQmlApplicationEngine>
#include <QString>
#include <QMap>

#include "EventEmitter.h"

typedef void (*GoCallbackFunc)(const char*, const char*, const char*);

class CuteGo
{
public:
    CuteGo(GoCallbackFunc gocallback);
    void start();
    void loadQmlEntry(QString qmlEntry);
    void newEventManager(QString name);

private:
    QGuiApplication *app;
    QQmlApplicationEngine *engine;
    QMap<QString, EventManager*> eventEmitters;
    GoCallbackFunc gc;
};

#endif // CUTEGO_H
