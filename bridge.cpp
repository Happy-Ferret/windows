#include "bridge.hpp"
#include "_cgo_export.h"
#include <iostream>


// Test to try c++ dll communication with go.
void CPP_OnPrint(const char *str) { goCallback((char *)str); }