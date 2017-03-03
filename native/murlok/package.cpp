#include "stdafx.h"

#include "package.h"

const std::wstring getPackageFamilyName(){
	UINT32 length = 0;
	LONG rc = GetCurrentPackageFamilyName(&length, NULL);
	if (rc != ERROR_INSUFFICIENT_BUFFER)
	{
		throw L"Process has no package identity\n";
	}

	PWSTR familyName = (PWSTR)malloc(length * sizeof(*familyName));
	if (familyName == NULL)
	{
		throw L"Error allocating memory\n";
	}

	rc = GetCurrentPackageFamilyName(&length, familyName);
	if (rc != ERROR_SUCCESS)
	{
		throw L"can't retrieve PackageFamilyName\n";
	}

	std::wstring ret(familyName);
	free(familyName);
	return ret;
}