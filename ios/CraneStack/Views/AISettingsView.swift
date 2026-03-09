import SwiftUI

struct AISettingsView: View {
    @EnvironmentObject var authManager: AppAuthManager

    var body: some View {
        List {
            Section("MCP Access Token") {
                Button("Generate MCP Token") {
                    authManager.generateMCPAccessToken()
                }

                if let token = authManager.mcpAccessToken {
                    HStack {
                        Text(token)
                            .font(.system(.caption, design: .monospaced))
                            .lineLimit(1)
                            .truncationMode(.middle)

                        Spacer()

                        Button {
                            UIPasteboard.general.string = token
                        } label: {
                            Image(systemName: "doc.on.doc")
                        }
                        .buttonStyle(.borderless)
                    }
                }

                if let error = authManager.mcpTokenError {
                    Text(error)
                        .font(.caption)
                        .foregroundStyle(.red)
                }
            }
        }
        .navigationTitle("AI Settings")
    }
}

#Preview {
    NavigationStack {
        AISettingsView()
            .environmentObject(AppAuthManager())
    }
}
