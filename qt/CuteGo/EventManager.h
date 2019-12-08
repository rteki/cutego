#ifndef EVENTEMITTER_H
#define EVENTEMITTER_H

#include <QObject>
#include <QVariant>
#include <QVector>
#include <QJSValue>
#include <QQmlApplicationEngine>

typedef void (*GoCallbackFunc)(const char*, const char*, const char*);


class EventManager : public QObject
{
    Q_OBJECT
public:
    explicit EventManager(QQmlApplicationEngine* _engine, QString _name, GoCallbackFunc gocallback, QObject *parent = nullptr);
    Q_INVOKABLE void call(QString eventName, QJSValue value);
    Q_INVOKABLE void on(QString eventName, QJSValue subscriber);

    void callQml(QString eventName, QString value);

    static QJSValue jsonParseFunc;
    static QJSValue jsonStringifyFunc;

signals:
    void signal_callQml(QString eventName, QString value);

public slots:
    void slot_callQml(QString eventName, QString value);

private:
    QString _name;
    GoCallbackFunc _gc;
    QMap<QString, QVector<QJSValue>*> _handlers;
    QQmlApplicationEngine* _engine;


};


#endif // EVENTEMITTER_H
