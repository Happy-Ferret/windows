#include "bridge.hpp"
#include <iostream>
#include "_cgo_export.h"

// Test to try c++ dll communication with go.
void CPP_OnPrint(const char* str){
     goCallback((char*)str);
}