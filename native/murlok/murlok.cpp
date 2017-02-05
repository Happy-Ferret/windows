#include "stdafx.h"

#include "murlok.h"


void InitOnMurlokPrint(FuncOnMurlokPrint fn) {
  Murlok_OnMurlokPrint = fn;
  Murlok_OnMurlokPrint("Hey Go do you hear me?");
}

void Driver_Run() {

}