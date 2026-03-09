import SwiftUI

struct MainTabView: View {
    @EnvironmentObject var authManager: AppAuthManager
    @StateObject private var subscriptionManager = AppSubscriptionManager()

    var body: some View {
        TabView {
            Text("Home")
                .tabItem {
                    Label("Home", systemImage: "house")
                }

            SubscriptionView()
                .environmentObject(authManager)
                .environmentObject(subscriptionManager)
                .tabItem {
                    Label("Subscription", systemImage: "creditcard")
                }

            SettingsView()
                .environmentObject(authManager)
                .tabItem {
                    Label("Settings", systemImage: "gear")
                }
        }
    }
}

#Preview {
    MainTabView()
        .environmentObject(AppAuthManager())
}
