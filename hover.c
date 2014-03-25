#include "hover.h"

void HoverCreateMessageBox(char* title,char* body,int flag){
    //这儿还可以做很多事
    MessageBox(0,body,title,flag);
}