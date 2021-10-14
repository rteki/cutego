#include "CuteGo.h"

#include <QResource>
#include <QQmlContext>
#include <QDebug>

const QString DEFAULT_RESOURCES_RCC = "./resources.rcc";

QJSValue EventManager::jsonParseFunc;
QJSValue EventManager::jsonStringifyFunc;

CuteGo::CuteGo(GoCallbackFunc gocallback)
{
    int argc = 0;
    char ** argv;

    gc = gocallback;

    app = new QGuiApplication(argc, argv);

    engine = new QQmlApplicationEngine();



    EventManager::jsonParseFunc = engine->evaluate("(function(str){return JSON.parse(str);})");
    EventManager::jsonStringifyFunc = engine->evaluate("(function(obj){return JSON.stringify(obj);})");
}

void CuteGo::registerResource(QString path) {
    QResource::registerResource(path);
}

void CuteGo::loadQmlEntry(QString qmlEntry)
{
    engine->load(QUrl("qrc:///" + qmlEntry));
}

void CuteGo::newEventManager(QString name)
{
    EventManager * ee = new EventManager(this->engine, name, gc);

    if(eventEmitters[name]) {
        delete eventEmitters[name];
        qDebug() << "Ololo";
    }

    eventEmitters[name] = ee;

    engine->rootContext()->setContextProperty(name, ee);
}

void CuteGo::call(QString eventManagerName, QString eventName, QString value)
{
    EventManager * em = this->eventEmitters[eventManagerName];

    if(!em) {
        qCritical() << "Can't find event manager called: " << eventManagerName << '\n';
        return;
    }

    em->callQml(eventName, value);
}

void CuteGo::start()
{
    app->exec();
}




