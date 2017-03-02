#include "stdafx.h"

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
	cout << "2" << endl;
    Platform::String ^ packageFamilyName = package->Id->FamilyName;
	cout << "3" << endl;


    // Create and set the connection
    auto connection = ref new AppServiceConnection();
	cout << "4" << endl;

    connection->PackageFamilyName = packageFamilyName;
    connection->AppServiceName = "CommunicationService";
    cout << "opening bridge..." << endl;

    // Open the connection
    create_task(connection->OpenAsync())
        .then([connection](AppServiceConnectionStatus status) {
          if (status != AppServiceConnectionStatus::Success) {
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
                           AppServiceRequestReceivedEventArgs ^ args) {}

void BridgeClosed(AppServiceConnection ^ connection,
                  AppServiceClosedEventArgs ^ args) {}