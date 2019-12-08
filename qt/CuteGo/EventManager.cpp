#include <QDebug>
#include "EventManager.h"
#include <QJsonObject>
#include <QJsonDocument>
#include <QJSEngine>


#include <QDebug>

EventManager::EventManager(QQmlApplicationEngine* engine, QString name, GoCallbackFunc gocallback, QObject *parent) : QObject(parent)
{
    this->_engine = engine;
    this->_name = name;
    this->_gc = gocallback;

    connect(this, EventManager::signal_callQml, this, EventManager::slot_callQml);
}

void EventManager::call(QString eventName, QJSValue value)
{
    QJSValue jsJsonString = EventManager::jsonStringifyFunc.call({value});

    _gc(this->_name.toStdString().c_str(), eventName.toStdString().c_str(), jsJsonString.toString().toUtf8().toStdString().c_str());
}

void EventManager::on(QString eventName, QJSValue subscriber)
{


    if(!_handlers[eventName]) {
        this->_handlers[eventName] = new QVector<QJSValue>();
    }

    QVector<QJSValue>* subscribers = _handlers[eventName];

    for(auto s : *subscribers) {
        if(s.equals(subscriber)) {
            qWarning() << "CuteGo QML Warning: " << "Subscriber for '" + eventName + "\' is already registered";
            return;
        }
    }

    subscribers->push_back(subscriber);

}

void EventManager::callQml(QString eventName, QString value)
{
    emit signal_callQml(eventName, value);
}

void EventManager::slot_callQml(QString eventName, QString value)
{

    QVector<QJSValue>* subscribers = _handlers[eventName];

    if(!subscribers) {
        qWarning() << "CuteGo QML Warning: " << "Can't find handler for '" << eventName << '\'';
        return;
    }

    for(auto subscriber : *subscribers) {
        if(subscriber.isCallable()) {
            QJSValue parsedJSVal = EventManager::jsonParseFunc.call({value});
            subscriber.call({parsedJSVal});
        }
        else {
            qCritical() << "CuteGo QML Warning: " << "Can't call handler for '" << eventName << "\'. Is not callable!";
        }
    }
}
