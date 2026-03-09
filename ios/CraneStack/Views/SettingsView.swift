import SwiftUI
import CraneKit
import CraneKitUI

struct SettingsView: View {
    @EnvironmentObject var authManager: AuthManager
    @StateObject private var subscriptionManager = AppSubscriptionManager()

    var body: some View {
        NavigationStack {
            List {
                NavigationLink("Subscription") {
                    CraneKitUI.SubscriptionView(priceLabel: "$4.99 / month")
                        .environmentObject(subscriptionManager)
                }
                NavigationLink("AI Settings") {
                    MCPTokenView()
                        .environmentObject(authManager)
                }
            }
            .navigationTitle("Settings")
        }
    }
}

#Preview {
    SettingsView()
        .environmentObject(AuthManager())
}
