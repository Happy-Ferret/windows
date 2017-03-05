#include "stdafx.h"

#include "bridge.h"
#include "murlok.h"

using namespace std;

void Driver_Run()
{
  std::cout << "C++ ~> Driver_Run()" << std::endl;
  BridgeConnectAsync();

  string result;
  cin >> result;
}
