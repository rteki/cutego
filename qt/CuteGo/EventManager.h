#ifndef EVENTEMITTER_H
#define EVENTEMITTER_H

#include <QObject>
#include <QVariant>

typedef void (*GoCallbackFunc)(const char*, const char*, const char*);

class EventManager : public QObject
{
    Q_OBJECT
public:
    explicit EventManager(QString name, GoCallbackFunc gocallback, QObject *parent = nullptr);
    Q_INVOKABLE void call(QString eventName, QVariant param);

signals:

public slots:


private:
    QString name;
    GoCallbackFunc gc;
};

#endif // EVENTEMITTER_H
