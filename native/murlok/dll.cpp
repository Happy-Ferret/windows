#include "stdafx.h"
#include "dll.h"

 _onLaunch_t _onLaunch = nullptr;
 _onTerminate_t _onTerminate = nullptr;

void Init_OnLaunch(_onLaunch_t fn)
{
    _onLaunch = fn;
}

void Init_OnTerminate(_onTerminate_t fn)
{
    _onTerminate = fn;
}
