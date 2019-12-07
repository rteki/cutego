package tmpwriter

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type TmpWriter []string

func (w *TmpWriter) Write(path string, text string) {
	*w = append(*w, path)

	msg := []byte(text)

	err := ioutil.WriteFile(path, msg, 0644)

	if err != nil {
		fmt.Println(err)
	}

}

func (w *TmpWriter) Close() {
	for _, path := range *w {
		err := os.Remove(path)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (w *TmpWriter) ExecInDir(dir string, filename string, callback func() string) {
	prevdir, err := os.Getwd()

	if err != nil {
		fmt.Println(prevdir)
		fmt.Println(err)
		os.Exit(666)
	}

	os.Chdir(dir)

	qrc := callback()

	os.Chdir(prevdir)

	w.Write(path.Join(dir, filename), qrc)
}
