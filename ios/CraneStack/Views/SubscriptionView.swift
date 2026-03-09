import SwiftUI

struct SubscriptionView: View {
    @EnvironmentObject var authManager: AppAuthManager
    @EnvironmentObject var subscriptionManager: AppSubscriptionManager

    // TODO: Replace with a price fetched from your backend or StoreKit product.
    private let monthlyPriceLabel = "$4.99 / month"

    var body: some View {
        NavigationStack {
            List {
                Section("Monthly Plan") {
                    if subscriptionManager.isSubscribed {
                        HStack {
                            Image(systemName: "checkmark.seal.fill")
                                .foregroundStyle(.green)
                            Text("Subscribed")
                                .bold()
                        }

                        Button("Cancel Subscription", role: .destructive) {
                            // TODO: implement cancellation flow
                        }
                    } else {
                        VStack(alignment: .leading, spacing: 8) {
                            Text("Unlock full access with a monthly subscription.")
                                .font(.subheadline)
                                .foregroundStyle(.secondary)

                            Button {
                                Task {
                                    guard let token = authManager.sessionToken else { return }
                                    await subscriptionManager.subscribe(sessionToken: token)
                                }
                            } label: {
                                Text("Subscribe — \(monthlyPriceLabel)")
                                    .frame(maxWidth: .infinity)
                            }
                            .buttonStyle(.borderedProminent)
                        }
                        .padding(.vertical, 4)
                    }

                    if let error = subscriptionManager.subscriptionError {
                        Text(error)
                            .font(.caption)
                            .foregroundStyle(.red)
                    }
                }
            }
            .navigationTitle("Subscription")
            .task {
                guard let token = authManager.sessionToken else { return }
                await subscriptionManager.fetchStatus(sessionToken: token)
            }
        }
    }
}

#Preview {
    SubscriptionView()
        .environmentObject(AppAuthManager())
        .environmentObject(AppSubscriptionManager())
}
