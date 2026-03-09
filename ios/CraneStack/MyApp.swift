import SwiftUI
import CraneKit

@main
struct MyApp: App {
    init() {
        APIClient.configure(baseURL: URL(string: "http://localhost:8080")!)
    }

    @StateObject private var authManager = AuthManager(apiClient: .shared)

    var body: some Scene {
        WindowGroup {
            if authManager.isAuthenticated {
                MainTabView()
                    .environmentObject(authManager)
            } else {
                LoginView()
                    .environmentObject(authManager)
            }
        }
    }
}
