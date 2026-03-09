import SwiftUI

struct SettingsView: View {
    @EnvironmentObject var authManager: AppAuthManager
    @StateObject private var subscriptionManager = AppSubscriptionManager()

    var body: some View {
        NavigationStack {
            List {
                NavigationLink("Subscription") {
                    SubscriptionView()
                        .environmentObject(subscriptionManager)
                }
                NavigationLink("AI Settings") {
                    AISettingsView()
                }
            }
            .navigationTitle("Settings")
        }
    }
}

#Preview {
    SettingsView()
        .environmentObject(AppAuthManager())
}
