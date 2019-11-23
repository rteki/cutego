#include "CuteGo.h"

#include <QResource>
#include <QQmlContext>
#include <QDebug>

const QString DEFAULT_RESOURCES_RCC = "./resources.rcc";

CuteGo::CuteGo(GoCallbackFunc gocallback)
{
    int argc = 0;
    char ** argv;

    gc = gocallback;

    app = new QGuiApplication(argc, argv);

    engine = new QQmlApplicationEngine();

    QResource::registerResource(DEFAULT_RESOURCES_RCC);
}

void CuteGo::loadQmlEntry(QString qmlEntry)
{
    engine->load(QUrl("qrc:///" + qmlEntry));
}

void CuteGo::newEventManager(QString name)
{
    EventManager * ee = new EventManager(name, gc);

    if(eventEmitters[name]) {
        delete eventEmitters[name];
        qDebug() << "Ololo";
    }

    eventEmitters[name] = ee;

    engine->rootContext()->setContextProperty(name, ee);
}

void CuteGo::start()
{
    app->exec();
}




