#include "stdafx.h"

#include "dll.h"
#include "bridge.h"

using namespace std;
using namespace concurrency;
using namespace Platform;
using namespace Windows::ApplicationModel::AppService;
using namespace Windows::ApplicationModel::DataTransfer;
using namespace Windows::Foundation;
using namespace Windows::Foundation::Collections;
using namespace Windows::Storage;
using namespace Windows::UI::Notifications;
using namespace Windows::Data::Xml::Dom;

AppServiceConnection ^ bridgeConn = nullptr;

IAsyncAction ^ BridgeConnectAsync() {
    return create_async([] {
        // Get the package family name
        Windows::ApplicationModel::Package ^ package =
            Windows::ApplicationModel::Package::Current;
        Platform::String ^ packageFamilyName = package->Id->FamilyName;

        // Create and set the connection
        auto connection = ref new AppServiceConnection();

        connection->PackageFamilyName = packageFamilyName;
        connection->AppServiceName = "CommunicationService";
        cout << "opening bridge..." << endl;

        // Open the connection
        create_task(connection->OpenAsync())
            .then([connection](AppServiceConnectionStatus status) {
                if (status != AppServiceConnectionStatus::Success)
                {
                    cout << "bridge connection failed: " << (int)status << endl;
                    return;
                }

                bridgeConn = connection;
                cout << "bridge ready" << endl;

                bridgeConn->RequestReceived +=
                    ref new TypedEventHandler<AppServiceConnection ^,
                                              AppServiceRequestReceivedEventArgs ^>(
                        BridgeRequestReceived);
                bridgeConn->ServiceClosed +=
                    ref new TypedEventHandler<AppServiceConnection ^,
                                              AppServiceClosedEventArgs ^>(
                        BridgeClosed);
            });
    });
};

void BridgeRequestReceived(AppServiceConnection ^ connection,
                           AppServiceRequestReceivedEventArgs ^ args)
{
    auto deferral = args->GetDeferral();

    auto message = args->Request->Message;
    auto actionType = _wtoi(message->Lookup(L"type")->ToString()->Data());
    auto actionPayload = message->Lookup(L"payload")->ToString();

    cout << "action type: " << actionType << endl;

    switch (actionType)
    {
    case DriverLaunched:
		cout << "DriverLaunched " << endl;
		_onLaunch();
		cout << "DriverLaunched ok" << endl;
        break;

    default:
        break;
    }

    auto response = ref new ValueSet();
    response->Insert("response", "ack");
    create_task(args->Request->SendResponseAsync(response)).then([deferral](AppServiceResponseStatus status) {
        deferral->Complete();
    });
}

void BridgeClosed(AppServiceConnection ^ connection,
                  AppServiceClosedEventArgs ^ args)
{
    cout << "bridge closed" << endl;
}