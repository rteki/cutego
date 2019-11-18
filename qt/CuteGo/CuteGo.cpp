#include "CuteGo.h"

#include <QResource>
#include <QtGui/QGuiApplication>
#include <QtQml/QQmlApplicationEngine>

CuteGo::CuteGo()
{
}

void CuteGo::setResourcesPath(QString path, QString entry)
{
    this->resourcesPath = path;
    this->entry = entry;

    QResource::registerResource(path);

}

void CuteGo::start()
{
    int argc;
    char ** argv;

    QGuiApplication app(argc, argv);
    QQmlApplicationEngine engine;
    engine.load(QUrl(QStringLiteral("qrc:///qml/main.qml")));

    app.exec();
}


