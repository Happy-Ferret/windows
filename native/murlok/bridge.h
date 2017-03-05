// bridge.h
#pragma once

using namespace Windows::ApplicationModel::AppService;
using namespace Windows::Foundation;

enum ActionType
{
	DriverLaunched,
	DriverTerminating,
	DriverTerminated
};

IAsyncAction ^ BridgeConnectAsync();
void BridgeRequestReceived(AppServiceConnection ^ connection,
                           AppServiceRequestReceivedEventArgs ^ args);
void BridgeClosed(AppServiceConnection ^ connection,
                  AppServiceClosedEventArgs ^ args);
