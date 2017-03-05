// dll.h
#pragma once

typedef unsigned int (*_onLaunch_t)();
typedef unsigned int (*_onTerminate_t)();

 extern _onLaunch_t _onLaunch;
 extern _onTerminate_t _onTerminate;

extern "C" {
__declspec(dllexport) void Init_OnLaunch(_onLaunch_t fn);
__declspec(dllexport) void Init_OnTerminate(_onTerminate_t fn);
}
