#include "stdafx.h"
#include "dll.h"

_voidCallback_t _onLaunch = nullptr;
_voidCallback_t _onFocus = nullptr;
_voidCallback_t _onBlur = nullptr;
_onTerminate_t _onTerminate = nullptr;
_voidCallback_t _onFinalize = nullptr;

void Init_OnLaunch(_voidCallback_t fn)
{
    _onLaunch = fn;
}

void Init_OnFocus(_voidCallback_t fn)
{
	_onFocus = fn;
}

void Init_OnBlur(_voidCallback_t fn)
{
	_onBlur = fn;
}

void Init_OnTerminate(_onTerminate_t fn)
{
    _onTerminate = fn;
}

void Init_OnFinalize(_voidCallback_t fn)
{
    _onFinalize = fn;
}