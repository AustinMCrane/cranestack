import CraneKit
import AuthenticationServices
import Foundation

/// App-specific auth manager. Inherits session state, devLogin, and APIClient
/// from CraneKit's AuthManager base class.
@MainActor
final class AppAuthManager: AuthManager {
    @Published var mcpAccessToken: String?
    @Published var mcpTokenError: String?

    /// Signs the user in with Apple ID.
    func loginWithApple() {
        // TODO: create ASAuthorizationAppleIDRequest
        // TODO: present ASAuthorizationController
        // TODO: on success, call completeLogin(token:) with the session token
    }

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
