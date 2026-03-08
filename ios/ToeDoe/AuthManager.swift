import AuthenticationServices
import Combine
import Foundation

/// Manages the user's authentication state and token lifecycle.
@MainActor
final class AuthManager: ObservableObject {
    @Published var isAuthenticated: Bool = false
    @Published var mcpAccessToken: String?
    @Published var mcpTokenError: String?

    let apiClient = APIClient()

    /// Signs the user in with Apple ID.
    func loginWithApple() {
        // TODO: create ASAuthorizationAppleIDRequest
        // TODO: present ASAuthorizationController
        // TODO: on success, send the identity token to POST /auth/login
        // TODO: store the returned session token via apiClient.setAuthToken(_:) and set isAuthenticated = true
    }

    #if DEBUG
    /// Bypasses authentication for local development. Remove before shipping.
    /// Start the API server with DEV_SESSION_TOKEN=dev-session-token to match.
    func devLogin() {
        apiClient.setAuthToken("dev-session-token")
        isAuthenticated = true
    }
    #endif

    /// Requests a new Personal Access Token (PAT) for use with the MCP server.
    func generateMCPAccessToken() {
        mcpTokenError = nil
        Task {
            do {
                let data = try await apiClient.post("/auth/generate-mcp-key", body: EmptyBody())
                let response = try JSONDecoder().decode(GenerateMCPKeyResponse.self, from: data)
                mcpAccessToken = response.token
            } catch {
                mcpTokenError = error.localizedDescription
                print("generateMCPAccessToken error: \(error)")
            }
        }
    }
}

// MARK: - Response types

private struct EmptyBody: Encodable {}

private struct GenerateMCPKeyResponse: Decodable {
    let token: String
}
