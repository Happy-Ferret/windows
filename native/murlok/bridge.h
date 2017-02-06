#pragma once

using namespace Windows::ApplicationModel::AppService;
using namespace Windows::Foundation;

IAsyncAction ^ BridgeConnectAsync();
void BridgeRequestReceived(AppServiceConnection ^ connection,
                           AppServiceRequestReceivedEventArgs ^ args);
void BridgeClosed(AppServiceConnection ^ connection,
                  AppServiceClosedEventArgs ^ args);
