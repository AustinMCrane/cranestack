import SwiftUI
import CraneKit
import CraneKitUI

struct LoginView: View {
    var body: some View {
        AppleSignInView {
            Text("My App")
                .font(.largeTitle)
                .bold()
        }
    }
}

#Preview {
    LoginView()
        .environmentObject(AuthManager())
}
