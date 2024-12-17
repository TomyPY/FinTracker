import '../styles/globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { GoogleOAuthProvider } from "@react-oauth/google"

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Auth App',
  description: 'Una aplicación de autenticación simple',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="es" className={inter.className}>
      <GoogleOAuthProvider clientId='1009282809407-sh8h2kgmot2q295a503sl5530pldnaj9.apps.googleusercontent.com'>
        <body>{children}</body>
      </GoogleOAuthProvider>
    </html>
  )
}

