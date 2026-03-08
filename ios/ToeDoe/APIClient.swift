import Foundation

enum APIError: Error {
    case missingAuthToken
    case httpError(Int)
}

/// Handles authenticated HTTP requests to the ToeDoe REST API.
final class APIClient {
    private let baseURL: URL
    private var authToken: String?

    init(baseURL: URL = URL(string: "http://localhost:8080")!) {
        self.baseURL = baseURL
    }

    /// Attaches the session token used for authenticated requests.
    func setAuthToken(_ token: String) {
        // TODO: store token securely in Keychain
        authToken = token
    }

    /// Performs an authenticated GET request to the given path.
    func get(_ path: String) async throws -> Data {
        let request = try buildRequest(path: path, method: "GET", body: nil as Data?)
        return try await execute(request)
    }

    /// Performs an authenticated POST request with a JSON-encodable body.
    func post<Body: Encodable>(_ path: String, body: Body) async throws -> Data {
        let bodyData = try JSONEncoder().encode(body)
        var request = try buildRequest(path: path, method: "POST", body: bodyData)
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        return try await execute(request)
    }

    // MARK: - Private

    private func buildRequest(path: String, method: String, body: Data?) throws -> URLRequest {
        guard let token = authToken else { throw APIError.missingAuthToken }
        let url = baseURL.appending(path: path)
        var request = URLRequest(url: url)
        request.httpMethod = method
        request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        request.httpBody = body
        return request
    }

    private func execute(_ request: URLRequest) async throws -> Data {
        let (data, response) = try await URLSession.shared.data(for: request)
        if let http = response as? HTTPURLResponse, !(200...299).contains(http.statusCode) {
            throw APIError.httpError(http.statusCode)
        }
        return data
    }
}
