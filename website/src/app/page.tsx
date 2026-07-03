import Link from "next/link";

export default function Home() {
  return (
    <main className="min-h-screen">
      {/* Header */}
      <header className="sticky top-0 z-50 border-b border-border bg-white/80 backdrop-blur-md">
        <div className="max-w-5xl mx-auto px-6 h-14 flex items-center justify-between">
          <Link href="/" className="flex items-center gap-2 font-bold text-lg">
            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/>
            </svg>
            Velo
          </Link>
          <nav className="flex items-center gap-6 text-sm text-muted-foreground">
            <a href="#features" className="hover:text-foreground transition">Features</a>
            <a href="#benchmark" className="hover:text-foreground transition">Benchmark</a>
            <a href="https://github.com/MTconnect-BR/APIs" target="_blank" rel="noopener noreferrer" className="hover:text-foreground transition">
              GitHub
            </a>
          </nav>
        </div>
      </header>

      {/* Hero */}
      <section className="py-32 px-6">
        <div className="max-w-5xl mx-auto text-center">
          <h1 className="text-6xl md:text-8xl font-bold tracking-tight leading-none">
            Fast
          </h1>
          <h1 className="text-6xl md:text-8xl font-bold tracking-tight leading-none">
            Smart
          </h1>
          <h1 className="text-6xl md:text-8xl font-bold tracking-tight leading-none">
            Global
          </h1>
          <h1 className="text-6xl md:text-8xl font-bold tracking-tight leading-none text-primary">
            Velo&gt;API
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
                <div className="bg-[#F5F5F5] rounded-lg p-4 font-mono text-sm overflow-x-auto">
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

      {/* Features */}
      <section id="features" className="py-24 px-6">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold text-center mb-4">7 problems every API faces</h2>
          <p className="text-center text-muted-foreground mb-16">Velo solves them all.</p>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
            <FeatureCard icon="⚡" title="Rate Limiting" description="Token bucket + sliding window" />
            <FeatureCard icon="📦" title="Cache" description="Redis-backed with TTL" />
            <FeatureCard icon="🔐" title="Auth" description="JWT + OAuth2 + API keys" />
            <FeatureCard icon="⚖️" title="Load Balancing" description="Round-robin, least connections" />
            <FeatureCard icon="📊" title="Observability" description="Prometheus + structured logs" />
            <FeatureCard icon="📝" title="Docs" description="Auto OpenAPI 3.1" />
            <FeatureCard icon="🏷️" title="Versioning" description="Header + path + query" />
          </div>
        </div>
      </section>

      {/* Code Example */}
      <section className="py-24 px-6 bg-[#F9F9F9]">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold text-center mb-4">Simple configuration</h2>
          <p className="text-center text-muted-foreground mb-16">One YAML file. That&apos;s it.</p>
          <div className="max-w-2xl mx-auto">
            <div className="bg-white border border-border rounded-lg overflow-hidden">
              <div className="flex items-center gap-2 px-4 py-3 border-b border-border">
                <div className="w-3 h-3 rounded-full bg-[#FF5F56]" />
                <div className="w-3 h-3 rounded-full bg-[#FFBD2E]" />
                <div className="w-3 h-3 rounded-full bg-[#27C93F]" />
                <span className="ml-2 text-sm text-muted-foreground">velo.yaml</span>
              </div>
              <pre className="p-6 text-sm overflow-x-auto">
{`server:
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
      secret: "\${JWT_SECRET}"`}
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Benchmark */}
      <section id="benchmark" className="py-24 px-6">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold text-center mb-4">Performance</h2>
          <p className="text-center text-muted-foreground mb-16">Traditional API vs Velo Gateway</p>
          <div className="max-w-2xl mx-auto space-y-8">
            <BenchmarkRow label="Latência" traditional="45ms" velo="12ms" percent={27} />
            <BenchmarkRow label="Throughput" traditional="1k req/s" velo="50k req/s" percent={100} />
            <BenchmarkRow label="Memória" traditional="256 MB" velo="45 MB" percent={18} />
            <BenchmarkRow label="Cold Start" traditional="2.5s" velo="15ms" percent={1} />
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="py-24 px-6 bg-[#F9F9F9]">
        <div className="max-w-3xl mx-auto text-center">
          <h2 className="text-4xl font-bold mb-6">Ready to accelerate?</h2>
          <p className="text-muted-foreground mb-10">
            Deploy in minutes. Single binary. Zero dependencies.
          </p>
          <div className="flex gap-4 justify-center">
            <a href="https://github.com/MTconnect-BR/APIs" target="_blank" rel="noopener noreferrer" className="bg-primary text-white px-8 py-3 rounded-lg font-medium hover:bg-primary/90 transition">
              Get Started
            </a>
            <a href="https://github.com/MTconnect-BR/APIs" target="_blank" rel="noopener noreferrer" className="border border-border px-8 py-3 rounded-lg font-medium hover:bg-muted transition">
              View on GitHub
            </a>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-8 px-6 border-t border-border">
        <div className="max-w-5xl mx-auto flex items-center justify-between text-sm text-muted-foreground">
          <span>© 2024 Velo API</span>
          <div className="flex gap-6">
            <a href="https://github.com/MTconnect-BR/APIs" className="hover:text-foreground transition">GitHub</a>
            <span>MIT License</span>
          </div>
        </div>
      </footer>
    </main>
  );
}

function FeatureCard({ icon, title, description }: { icon: string; title: string; description: string }) {
  return (
    <div className="border border-border rounded-lg p-4 hover:shadow-sm transition">
      <div className="text-2xl mb-2">{icon}</div>
      <h3 className="font-semibold text-sm">{title}</h3>
      <p className="text-xs text-muted-foreground">{description}</p>
    </div>
  );
}

function BenchmarkRow({ label, traditional, velo, percent }: { label: string; traditional: string; velo: string; percent: number }) {
  return (
    <div>
      <div className="flex justify-between items-center mb-2">
        <span className="font-medium">{label}</span>
        <div className="text-sm">
          <span className="text-muted-foreground">{traditional}</span>
          <span className="mx-2">→</span>
          <span className="text-primary font-medium">{velo}</span>
        </div>
      </div>
      <div className="h-2 bg-muted rounded-full overflow-hidden">
        <div className="h-full bg-primary rounded-full" style={{ width: `${percent}%` }} />
      </div>
    </div>
  );
}
