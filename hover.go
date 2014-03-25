package hover

/*
#cgo LDFLAGS: -lhover
#include <windows.h>
#include "hover.h"
*/
import "C"

func Msgbox(title string, body string) int {
	C.MessageBox(nil, (*C.CHAR)(C.CString(body)), (*C.CHAR)(C.CString(title)), 0)
	return 0
}

func CreateMsgbox(title string, body string, flag int) {
	C.HoverCreateMessageBox(C.CString(body), C.CString(title), C.int(flag))
}
