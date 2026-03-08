import CraneKit
import Foundation

// APIClient is re-exported from CraneKit for use across the app.
// Configure the base URL here if it differs from the default (localhost:8080).
extension APIClient {
    static let shared = APIClient(baseURL: URL(string: "http://localhost:8080")!)
}
