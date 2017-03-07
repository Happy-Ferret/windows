#ifndef dll_h
#define dll_h

#ifdef __cplusplus
extern "C" {
#endif

void DLL_OnLaunch();
void DLL_OnFocus();
void DLL_OnBlur();
int DLL_OnTerminate();
void DLL_OnFinalize();

#ifdef __cplusplus
}
#endif

#endif /* dll_h */