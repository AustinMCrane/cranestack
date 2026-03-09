import CraneKit
import StoreKit

/// App-specific subscription manager wrapping CraneKit's StoreKit-based SubscriptionManager.
@MainActor
final class AppSubscriptionManager: ObservableObject {
    private let manager: SubscriptionManager

    @Published var isSubscribed: Bool = false
    @Published var subscriptionError: String?

    /// Pass your App Store product identifiers here.
    init(productIDs: [String] = []) {
        manager = SubscriptionManager(productIDs: productIDs)
    }

    /// Refreshes subscription status from StoreKit entitlements.
    func fetchStatus() async {
        await manager.refreshStatus()
        if case .subscribed = manager.status {
            isSubscribed = true
        } else {
            isSubscribed = false
        }
    }

    /// Purchases the first available subscription product via StoreKit.
    func subscribe() async {
        guard let product = manager.products.first else {
            subscriptionError = "No subscription products available."
            return
        }
        do {
            try await manager.purchase(product)
            if case .subscribed = manager.status {
                isSubscribed = true
            }
        } catch {
            subscriptionError = error.localizedDescription
        }
    }
}
