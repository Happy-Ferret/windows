// This is the main DLL file.

#include "stdafx.h"

#include "murlok.h"
#include <iostream>

void InitOnMurlokPrint(FuncOnMurlokPrint fn) {
	Murlok_OnMurlokPrint = fn;
	Murlok_OnMurlokPrint("Hey Go do you hear me?");
}