// murlok.h

#pragma once
#include <cstddef>

typedef void (*FuncOnMurlokPrint)(const char *);

FuncOnMurlokPrint Murlok_OnMurlokPrint = NULL;

extern "C" __declspec(dllexport) void InitOnMurlokPrint(FuncOnMurlokPrint fn);
extern "C" __declspec(dllexport) void Driver_Run();
