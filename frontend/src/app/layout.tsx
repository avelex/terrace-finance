import type { Metadata } from "next";
import { Source_Sans_3 } from "next/font/google";
import "./globals.css";
import { WalletProvider } from "@/contexts/WalletContext";

const sourceSans = Source_Sans_3({
  variable: "--font-source-sans",
  subsets: ["latin"],
  weight: ["400", "600", "700"],
});

export const metadata: Metadata = {
  title: "Terrace Finance",
  description: "DeFi yield protocol",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${sourceSans.variable} antialiased`}
      >
        <div className="app-background"></div>
        <WalletProvider>
          {children}
        </WalletProvider>
      </body>
    </html>
  );
}
