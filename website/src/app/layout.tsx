import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Velo - API Gateway",
  description: "High-performance API gateway for modern applications",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="antialiased">
        {children}
      </body>
    </html>
  );
}
