import SwiftUI

struct MainTabView: View {
    @EnvironmentObject var authManager: AppAuthManager

    var body: some View {
        TabView {
            Text("Home")
                .tabItem {
                    Label("Home", systemImage: "house")
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
