#include "stdafx.h"

#include "bridge.h"
#include "murlok.h"

void InitOnMurlokPrint(FuncOnMurlokPrint fn) {
  Murlok_OnMurlokPrint = fn;
  Murlok_OnMurlokPrint("Hey Go do you hear me?");
}

void Driver_Run() { 
	std::cout << "C++ ~> Driver_Run()" << std::endl;
	Windows::Foundation::Initialize();
	BridgeConnectAsync();
}

void HelloGo() {
	std::cout << "Hello Go, I'm cpp :D" << std::endl;
}