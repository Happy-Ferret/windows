
#include "dll.hpp"
#include "_cgo_export.h"
#include <iostream>

void DLL_OnLaunch()
{
    onLaunch();
}

void DLL_OnFocus()
{
    onFocus();
}

void DLL_OnBlur()
{
    onBlur();
}

int DLL_OnTerminate()
{
    return onTerminate();
}

void DLL_OnFinalize()
{
    onFinalize();
}