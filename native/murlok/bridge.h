// bridge.h
#pragma once

using namespace Windows::ApplicationModel::AppService;
using namespace Windows::Foundation;

enum ActionType
{
    OnLaunch,
	OnFocus,
	OnBlur,
    OnTerminate,
    OnFinalize
};

IAsyncAction ^ BridgeConnectAsync();
void BridgeRequestReceived(AppServiceConnection ^ connection,
			   AppServiceRequestReceivedEventArgs ^ args);
void BridgeClosed(AppServiceConnection ^ connection,
		  AppServiceClosedEventArgs ^ args);
