# Plan: Move Boilerplate from CraneStack iOS → CraneKit

## Problem
The CraneStack iOS app contains significant boilerplate that every CraneKit-based app would need
identically. Moving this to CraneKit makes it the one-stop SDK: drop it in, configure a URL,
and get auth, subscriptions, and common UI out of the box.

## Approach
Move each piece of boilerplate into CraneKit, making it public and configurable. The app then
slims down to only truly app-specific code (tab structure, branding, business logic).

---

## What to Move

### 1. `APIClient.shared` — configurable singleton (CraneKit)
- Add `APIClient.configure(baseURL:)` + `APIClient.shared` to CraneKit
- App just calls `APIClient.configure(baseURL: ...)` at startup
- **Deletes:** `ios/CraneStack/APIClient.swift`

### 2. `AppAuthManager` — delete the wrapper (CraneKit already covers it)
- CraneKit's `AuthManager` is already `open` and does everything
- App creates `@StateObject private var authManager = AuthManager(apiClient: .shared)`
- **Deletes:** `ios/CraneStack/AuthManager.swift`

### 3. `AppSubscriptionManager` — move wrapper to CraneKit
- The `isSubscribed`/`subscribe()`/`fetchStatus()` wrapper over raw `SubscriptionManager` is universal
- Add it to CraneKit as a public class (keep name `AppSubscriptionManager` or rename)
- **Deletes:** `ios/CraneStack/SubscriptionManager.swift`

### 4. `AppleSignInView` — reusable login view (CraneKitUI)
- The Apple Sign-In button wiring + dev login button is identical in every app
- Add to CraneKitUI as `AppleSignInView` with a `@ViewBuilder header` slot for branding
- App's `LoginView` becomes ~5 lines wrapping `AppleSignInView` with a title
- CraneKit needs `AuthenticationServices` dependency

### 5. `MCPTokenView` — MCP token UI (CraneKitUI)
- `AISettingsView` content is 100% reusable: generate button, display token, copy button, error
- Add to CraneKitUI as `MCPTokenView`
- App's `AISettingsView` becomes a 1-line `NavigationLink` destination
- **Deletes:** `ios/CraneStack/Views/AISettingsView.swift`

### 6. `SubscriptionView` — configurable subscription UI (CraneKitUI)
- The subscription status/purchase UI is generic; only price label + copy are app-specific
- Add to CraneKitUI as `SubscriptionView(priceLabel:)` with customizable strings
- **Deletes:** `ios/CraneStack/Views/SubscriptionView.swift`

---

## What Stays in the App
- `MyApp.swift` — `@main` entry point (Swift can't subclass this), but slimmed down
- `MainTabView.swift` — app-specific tab structure
- `SettingsView.swift` — app-specific settings list (links to CraneKit views)
- `LoginView.swift` — thin wrapper around `AppleSignInView` with app name/branding

---

## CraneKit Package Changes
- Add a new `CraneKitUI` library target in `Package.swift` for SwiftUI views
- Keep core `CraneKit` lean (no SwiftUI dependency)
- Views guarded with `#if canImport(SwiftUI)` and platform checks (iOS/macOS)
- Follow same pattern as existing `#if canImport(StoreKit)` guard in `SubscriptionManager`

---

## Implementation Order
1. Add `APIClient.shared` singleton config to CraneKit
2. Delete `AppAuthManager` wrapper; update app to use `AuthManager` directly
3. Move `AppSubscriptionManager` wrapper to CraneKit
4. Add `CraneKitUI` target to `Package.swift` with SwiftUI views
5. Add `AppleSignInView` to CraneKitUI
6. Add `MCPTokenView` to CraneKitUI
7. Add `SubscriptionView` to CraneKitUI
8. Update app to import and use all CraneKitUI views; delete empty boilerplate files
