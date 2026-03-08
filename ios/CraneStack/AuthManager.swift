import CraneKit
import AuthenticationServices

/// App-specific auth manager.
///
/// loginWithApple(authorization:), generateMCPAccessToken(), devLogin(),
/// mcpAccessToken, and mcpTokenError are all provided by CraneKit's AuthManager
/// base class — no app-level boilerplate needed.
@MainActor
final class AppAuthManager: AuthManager {
    override init(apiClient: APIClient = .shared) {
        super.init(apiClient: apiClient)
    }
}
