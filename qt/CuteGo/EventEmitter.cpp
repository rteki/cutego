#include <QDebug>
#include "EventEmitter.h"


EventManager::EventManager(QString name, GoCallbackFunc gocallback, QObject *parent) : QObject(parent)
{
    this->name = name;
    this->gc = gocallback;
}

void EventManager::call(QString eventName, QString param)
{
    gc(this->name.toStdString().c_str(), eventName.toStdString().c_str(), param.toStdString().c_str());
}
