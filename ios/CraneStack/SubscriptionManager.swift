import CraneKit

/// App-specific subscription manager.
///
/// fetchStatus(sessionToken:) and subscribe(sessionToken:) are provided by
/// CraneKit's SubscriptionManager base class.
@MainActor
final class AppSubscriptionManager: SubscriptionManager {
    override init(apiClient: APIClient = .shared) {
        super.init(apiClient: apiClient)
    }
}
