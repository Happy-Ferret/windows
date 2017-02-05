// murlok.h

#pragma once
#include <cstddef>

using namespace System;

namespace murlok {

	public ref class Class1
	{
	};
}

typedef void(*FuncOnMurlokPrint)(const char*);

FuncOnMurlokPrint Murlok_OnMurlokPrint = NULL;

extern "C" __declspec(dllexport) void InitOnMurlokPrint(FuncOnMurlokPrint fn);