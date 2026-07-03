import Link from "next/link";

export default function Home() {
  return (
    <main className="min-h-screen">
      {/* Header */}
      <header className="sticky top-0 z-50 border-b border-border bg-background/80 backdrop-blur-md">
        <div className="max-w-5xl mx-auto px-6 h-14 flex items-center justify-between">
          <Link href="/" className="flex items-center gap-2 font-bold text-lg">
            Velo
          </Link>
          <nav className="flex items-center gap-6 text-sm text-muted-foreground">
            <a href="#docs" className="hover:text-foreground transition">Docs</a>
            <a href="#tests" className="hover:text-foreground transition">Tests</a>
            <a href="#compare" className="hover:text-foreground transition">Compare</a>
            <Link href="/playground" className="hover:text-foreground transition">Playground</Link>
            <a href="https://github.com/MTconnect-BR/APIs" target="_blank" rel="noopener noreferrer" className="hover:text-foreground transition">
              GitHub
            </a>
          </nav>
        </div>
      </header>

      {/* Hero */}
      <section className="flex flex-col justify-center items-center w-full h-screen min-h-[680px] px-[5%]">
        <div className="text-center">
          <h1 className="text-[65px] md:text-[105px] font-bold uppercase leading-none tracking-tight">
            FAST
          </h1>
          <h1 className="text-[65px] md:text-[105px] font-bold uppercase leading-none tracking-tight">
            SMART
          </h1>
          <h1 className="text-[65px] md:text-[105px] font-bold uppercase leading-none tracking-tight">
            GLOBAL
          </h1>
          <h1 className="text-[65px] md:text-[105px] font-bold uppercase leading-none tracking-tight">
            <span className="text-muted-foreground">VELO&gt;</span>API
          </h1>
        </div>
      </section>

      {/* Summary */}
      <section className="py-16 px-6">
        <div className="max-w-5xl mx-auto">
          <div className="border border-border rounded-lg p-8">
            <div className="flex items-start justify-between flex-wrap gap-8">
              <div>
                <p className="text-sm text-muted-foreground">over</p>
                <p className="text-4xl font-bold">7,801,929,461</p>
                <p className="text-sm text-muted-foreground">requests served last 30d.</p>
              </div>
              <div className="flex-1 min-w-[300px]">
                <p className="text-muted-foreground mb-4">No build tools needed!</p>
                <div className="bg-[#272931] rounded-lg p-4 font-mono text-sm overflow-x-auto">
                  <code>
                    <span className="text-muted-foreground">import </span>
                    <span className="text-foreground">velo </span>
                    <span className="text-muted-foreground">from </span>
                    <span className="text-primary">&quot;https://velo.sh/velo@latest&quot;</span>
                  </code>
                </div>
                <div className="mt-4">
                  <button className="bg-primary text-white px-6 py-2 rounded-lg font-medium hover:bg-primary/90 transition">
                    Get Started
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Documentation */}
      <section id="docs" className="py-24 px-6">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold mb-4">Documentation</h2>
          <p className="text-muted-foreground mb-8">Get started with Velo in minutes.</p>

          <div className="space-y-6 max-w-3xl">
            <div>
              <h3 className="font-semibold mb-2">1. Import</h3>
              <div className="bg-[#272931] rounded-lg p-4 font-mono text-sm overflow-x-auto">
                <code>
                  <span className="text-muted-foreground">import </span>
                  <span className="text-foreground">velo </span>
                  <span className="text-muted-foreground">from </span>
                  <span className="text-primary">&quot;https://velo.sh/velo@latest&quot;</span>
                </code>
              </div>
            </div>

            <div>
              <h3 className="font-semibold mb-2">2. Configure</h3>
              <div className="bg-[#272931] rounded-lg p-4 font-mono text-sm overflow-x-auto">
                <pre>{`server:
  port: 8080

ratelimit:
  enabled: true
  default: "100/min"

cache:
  enabled: true
  driver: redis
  ttl: 5m

auth:
  enabled: true
  providers:
    - type: jwt
      secret: "\${JWT_SECRET}"`}</pre>
              </div>
            </div>

            <div>
              <h3 className="font-semibold mb-2">3. Run</h3>
              <div className="bg-[#272931] rounded-lg p-4 font-mono text-sm overflow-x-auto">
                <code>
                  <span className="text-foreground">./velo</span>
                  <span className="text-muted-foreground"> --config velo.yaml</span>
                </code>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Tests */}
      <section id="tests" className="py-24 px-6">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold mb-4">Tests</h2>
          <p className="text-muted-foreground mb-8">Performance benchmarks.</p>

          <div className="max-w-3xl space-y-8">
            <div>
              <div className="flex justify-between items-center mb-2">
                <span className="font-medium">Latency</span>
                <div className="text-sm">
                  <span className="text-muted-foreground">45ms</span>
                  <span className="mx-2">→</span>
                  <span className="text-primary font-medium">12ms</span>
                </div>
              </div>
              <div className="h-2 bg-muted rounded-full overflow-hidden">
                <div className="h-full bg-primary rounded-full" style={{ width: "27%" }} />
              </div>
            </div>

            <div>
              <div className="flex justify-between items-center mb-2">
                <span className="font-medium">Throughput</span>
                <div className="text-sm">
                  <span className="text-muted-foreground">1k req/s</span>
                  <span className="mx-2">→</span>
                  <span className="text-primary font-medium">50k req/s</span>
                </div>
              </div>
              <div className="h-2 bg-muted rounded-full overflow-hidden">
                <div className="h-full bg-primary rounded-full" style={{ width: "100%" }} />
              </div>
            </div>

            <div>
              <div className="flex justify-between items-center mb-2">
                <span className="font-medium">Memory</span>
                <div className="text-sm">
                  <span className="text-muted-foreground">256 MB</span>
                  <span className="mx-2">→</span>
                  <span className="text-primary font-medium">45 MB</span>
                </div>
              </div>
              <div className="h-2 bg-muted rounded-full overflow-hidden">
                <div className="h-full bg-primary rounded-full" style={{ width: "18%" }} />
              </div>
            </div>

            <div>
              <div className="flex justify-between items-center mb-2">
                <span className="font-medium">Cold Start</span>
                <div className="text-sm">
                  <span className="text-muted-foreground">2.5s</span>
                  <span className="mx-2">→</span>
                  <span className="text-primary font-medium">15ms</span>
                </div>
              </div>
              <div className="h-2 bg-muted rounded-full overflow-hidden">
                <div className="h-full bg-primary rounded-full" style={{ width: "1%" }} />
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Comparisons */}
      <section id="compare" className="py-24 px-6">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold mb-4">Compare</h2>
          <p className="text-muted-foreground mb-8">Velo vs traditional API gateways used by AI companies.</p>

          <div className="overflow-x-auto">
            <table className="w-full text-sm">
              <thead>
                <tr className="border-b border-border">
                  <th className="text-left py-3 px-4 font-medium text-muted-foreground">Feature</th>
                  <th className="text-left py-3 px-4 font-medium text-primary">Velo</th>
                  <th className="text-left py-3 px-4 font-medium text-muted-foreground">OpenAI</th>
                  <th className="text-left py-3 px-4 font-medium text-muted-foreground">Anthropic</th>
                  <th className="text-left py-3 px-4 font-medium text-muted-foreground">GitHub</th>
                  <th className="text-left py-3 px-4 font-medium text-muted-foreground">Gemini</th>
                </tr>
              </thead>
              <tbody>
                <tr className="border-b border-border">
                  <td className="py-3 px-4">Self-hosted</td>
                  <td className="py-3 px-4 text-primary">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                </tr>
                <tr className="border-b border-border">
                  <td className="py-3 px-4">Open Source</td>
                  <td className="py-3 px-4 text-primary">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                </tr>
                <tr className="border-b border-border">
                  <td className="py-3 px-4">Rate Limiting</td>
                  <td className="py-3 px-4 text-primary">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✓</td>
                </tr>
                <tr className="border-b border-border">
                  <td className="py-3 px-4">Cache</td>
                  <td className="py-3 px-4 text-primary">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                </tr>
                <tr className="border-b border-border">
                  <td className="py-3 px-4">Auth (JWT/OAuth2)</td>
                  <td className="py-3 px-4 text-primary">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">API Key</td>
                  <td className="py-3 px-4 text-muted-foreground">API Key</td>
                  <td className="py-3 px-4 text-muted-foreground">OAuth</td>
                  <td className="py-3 px-4 text-muted-foreground">API Key</td>
                </tr>
                <tr className="border-b border-border">
                  <td className="py-3 px-4">Load Balancing</td>
                  <td className="py-3 px-4 text-primary">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                  <td className="py-3 px-4 text-muted-foreground">✗</td>
                </tr>
                <tr className="border-b border-border">
                  <td className="py-3 px-4">Observability</td>
                  <td className="py-3 px-4 text-primary">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">Dashboard</td>
                  <td className="py-3 px-4 text-muted-foreground">Dashboard</td>
                  <td className="py-3 px-4 text-muted-foreground">Logs</td>
                  <td className="py-3 px-4 text-muted-foreground">Dashboard</td>
                </tr>
                <tr className="border-b border-border">
                  <td className="py-3 px-4">Auto Docs (OpenAPI)</td>
                  <td className="py-3 px-4 text-primary">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✓</td>
                  <td className="py-3 px-4 text-muted-foreground">✓</td>
                </tr>
                <tr>
                  <td className="py-3 px-4">Cost</td>
                  <td className="py-3 px-4 text-primary font-medium">Free</td>
                  <td className="py-3 px-4 text-muted-foreground">Pay per call</td>
                  <td className="py-3 px-4 text-muted-foreground">Pay per call</td>
                  <td className="py-3 px-4 text-muted-foreground">Pay per call</td>
                  <td className="py-3 px-4 text-muted-foreground">Pay per call</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-8 px-6 border-t border-border">
        <div className="max-w-5xl mx-auto flex items-center justify-between text-sm text-muted-foreground">
          <span>&copy; 2024 Velo API</span>
          <div className="flex gap-6">
            <a href="https://github.com/MTconnect-BR/APIs" className="hover:text-foreground transition">GitHub</a>
            <span>MIT License</span>
          </div>
        </div>
      </footer>
    </main>
  );
}
