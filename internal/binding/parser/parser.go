package parser

import (
	"encoding/xml"
	"fmt"
	"runtime"

	"github.com/therecipe/qt/internal/utils"
)

const (
	SIGNAL = "signal"
	SLOT   = "slot"

	IMPURE = "impure"
	PURE   = "pure"

	MOC         = "main"
	PLAIN       = "plain"
	CONSTRUCTOR = "constructor"
	DESTRUCTOR  = "destructor"

	CONNECT    = "Connect"
	DISCONNECT = "Disconnect"
	CALLBACK   = "callback"

	GETTER = "getter"
	SETTER = "setter"

	VOID = "void"
)

var (
	ClassMap        = make(map[string]*Class)
	SubnamespaceMap = make(map[string]bool)
)

func GetModule(s string) (m *Module) {

	switch runtime.GOOS {
	case "darwin", "linux":
		{
			xml.Unmarshal([]byte(utils.Load(fmt.Sprintf("/usr/local/Qt5.6.0/Docs/Qt-5.6/qt%v/qt%v.index", s, s))), &m)
		}

	case "windows":
		{
			xml.Unmarshal([]byte(utils.Load(fmt.Sprintf("C:\\Qt\\Qt5.6.0\\Docs\\Qt-5.6\\qt%v\\qt%v.index", s, s))), &m)
		}
	}

	m.Prepare()

	return m
}
