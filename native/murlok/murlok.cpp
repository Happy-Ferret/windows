#include "stdafx.h"

#include "bridge.h"
#include "murlok.h"

using namespace std;

void InitOnMurlokPrint(FuncOnMurlokPrint fn)
{
  Murlok_OnMurlokPrint = fn;
  Murlok_OnMurlokPrint("Hey Go do you hear me?");
}

void Driver_Run()
{
  std::cout << "C++ ~> Driver_Run()" << std::endl;
  BridgeConnectAsync();
}
