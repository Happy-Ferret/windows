// dll.h
#pragma once

typedef void (*_voidCallback_t)();
typedef int (*_onTerminate_t)();

extern _voidCallback_t _onLaunch;
extern _voidCallback_t _onFocus;
extern _voidCallback_t _onBlur;
extern _onTerminate_t _onTerminate;
extern _voidCallback_t _onFinalize;

extern "C" {
	__declspec(dllexport) void Init_OnLaunch(_voidCallback_t fn);
	__declspec(dllexport) void Init_OnFocus(_voidCallback_t fn);
	__declspec(dllexport) void Init_OnBlur(_voidCallback_t fn);
	__declspec(dllexport) void Init_OnTerminate(_onTerminate_t fn);
	__declspec(dllexport) void Init_OnFinalize(_voidCallback_t fn);
}
