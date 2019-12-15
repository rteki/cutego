import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQml 2.13



ApplicationWindow {
    id: applicationWindow
    height: 600
    width: 340

    maximumHeight: height
    minimumHeight: height

    maximumWidth: width
    minimumWidth: width

    Component.onCompleted: {
        show();
    }

    Column {
        id: column
        anchors.bottomMargin: 10
        anchors.topMargin: 10
        anchors.fill: parent
        spacing: 20

        Text {
            text: qsTr("<b>cutego</b> example project")
            font.family: "Courier"
            font.pixelSize: 17
        } //header Text

        Row {
            width: 300
            height: 59
            spacing: 10

            Text {
                text: qsTr("<b>Input:</b>")
                font.family: "Courier"
                horizontalAlignment: Text.AlignLeft
                verticalAlignment: Text.AlignVCenter
                font.pixelSize: 18
            }

            TextField {
                id: input
                text: qsTr("")

                property string prevText: ""

                onTextChanged: {
                    if(!/(^\d+$)|(^\d+\.\d*$)|(^$)/.test(text)) {
                        text = prevText;
                    }
                    else {
                        prevText = text;
                    }
                }
            }


        } //input Row

        GroupBox {
            width: 300
            height: 250
            font.pointSize: 14
            font.family: "Courier"
            title: qsTr("<b>Calculate function</b>")

            CheckBox {
                id: sqrtCheckBox
                x: 35
                y: 125
                text: qsTr("Sqrt")
            }

            CheckBox {
                id: cubeCheckBox
                x: 35
                y: 71
                text: qsTr("Cube")
            }

            CheckBox {
                id: sqrCheckBox
                x: 35
                y: 25
                text: qsTr("Sqr")
            }


        } //Functions GroupBox

        Button {
            y: 0
            text: qsTr("<b>Calculate</b>")
            font.family: "Courier"
            onClicked: {
                handler.call(
                            "calculate",
                            {
                                sqr: sqrCheckBox.checked,
                                cube: cubeCheckBox.checked,
                                sqrt: sqrtCheckBox.checked,
                                input: Number(input.text)
                            }
                            );
            }
        } // calculate Button

        Text {
            id: outputText
            text: qsTr("<b>Output:</b>")
            font.family: "Courier"
            font.pixelSize: 18
        }

        TextArea {
            id: output
            text: qsTr("")
            font.pointSize: 14
            font.family: "Courier"
            enabled: false
            color: "black"

            Component.onCompleted: {
                handler.on("show-result", onShowResult)
            }

            function onShowResult(result) {
                text = "";
                for(var func in result) {
                    text += func + "(input) = ";
                    text += result[func] + '\n'
                }
            }

        }//output TextArea


    } //main Column

}

