import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Velo - Fast, Smart, Global API Gateway",
  description: "A high-performance API gateway for modern applications. Rate limiting, cache, auth, load balancing, observability, docs, and versioning.",
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
