#include <QDebug>
#include "EventManager.h"
#include <QJsonObject>
#include <QJsonDocument>

#include <QDebug>

EventManager::EventManager(QString name, GoCallbackFunc gocallback, QObject *parent) : QObject(parent)
{
    this->name = name;
    this->gc = gocallback;
}

void EventManager::call(QString eventName, QVariant param)
{
    gc(this->name.toStdString().c_str(), eventName.toStdString().c_str(), QJsonDocument(QJsonObject::fromVariantMap(param.toMap())).toJson(QJsonDocument::Indented).toStdString().c_str());
}
