#ifndef CUTEGO_H
#define CUTEGO_H


#include <QString>


class CuteGo
{
public:
    CuteGo();
    void setResourcesPath(QString path, QString entry);
    void start();

private:
    QString resourcesPath;
    QString entry;
};

#endif // CUTEGO_H
