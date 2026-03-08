import AuthenticationServices
import SwiftUI

struct LoginView: View {
    @EnvironmentObject var authManager: AppAuthManager

    var body: some View {
        VStack(spacing: 24) {
            Text("ToeDoe")
                .font(.largeTitle)
                .bold()

            SignInWithAppleButton(.signIn) { request in
                request.requestedScopes = [.fullName, .email]
            } onCompletion: { result in
                switch result {
                case .success:
                    authManager.loginWithApple()
                case .failure(let error):
                    // TODO: surface error to user
                    print("Apple Sign-In failed: \(error)")
                }
            }
            .signInWithAppleButtonStyle(.black)
            .frame(height: 50)
            .padding(.horizontal, 40)

            #if DEBUG
            Button("Dev Login (skip auth)") {
                authManager.devLogin()
            }
            .font(.footnote)
            .foregroundStyle(.secondary)
            #endif
        }
    }
}

#Preview {
    LoginView()
        .environmentObject(AppAuthManager())
}
