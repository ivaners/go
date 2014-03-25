## Go语言调用C ##

下面就来实现一个自定义的函数库

    //hover.h
	#include <windows.h>
	void HoverCreateMessageBox(char* title,char* body,int flag);
	 
	//hover.c
	#include "hover.h"
	 
	void HoverCreateMessageBox(char* title,char* body,int flag){
	    //这儿还可以做很多事
	    MessageBox(0,body,title,flag);
	}

这个库导出HoverCreateMessageBox函数，供接下来的Go调用。
编译成动态链接库：

	gcc -c -o hover.o hover.c
	gcc -shared -o hover.dll hover.o -Wl,--out-implib,libhover.a

接下来我们要在上面实现的hover包中添加一个调用自己函数库的函数

	package hover
	//#cgo LDFLAGS: -lhover
	//#include <windows.h>
	//#include <hover.h>
	import "C"
	func Msgbox(title string,body string) int{
	    C.MessageBox(nil,(*C.CHAR)(C.CString(body)),(*C.CHAR)(C.CString(title)),0);
	    return 0;
	}
	 
	func CreateMsgbox(title string,body string,flag int){
	    C.HoverCreateMessageBox(C.CString(body),C.CString(title),C.int(flag));
	}


这个比刚刚的时候多了一些东西，首先是头部，加入了//#cgo LDFLAGS: -lhover 一行，这个是gcc的链接参数，也就是我们刚刚实现的库需要链接到生成的文件中，不然无法调用，这个跟C完全一样的理解，只是libhover.a需要放到include目录下，如果不是，需要指定CFLAG。然后多引入了hover.h头文件（我把它放到include所在的目录下了，如果不愿意放，可以添加//#cgo CFLAGS: -Imyincludedir ，这个作为gcc编译时的参数，也是C的知识。
后面添加了一个函数，调用刚刚库中的函数。

如果设置了GOPATH，可以把源码放在src包里，调用

	go get hover

在pkg目录下面会自动生成编译文件，最后，写一个测试程序来调用一下

	package main
 
	import "hover"
	 
	func main(){
	    hover.Msgbox("title","body");
	    hover.CreateMsgbox("title","body",64);//MB_ICONINFORMATION=64
	}