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
            <a href="#features" className="hover:text-foreground transition">Features</a>
            <a href="https://github.com/MTconnect-BR/APIs" target="_blank" rel="noopener noreferrer" className="hover:text-foreground transition">
              GitHub
            </a>
          </nav>
        </div>
      </header>

      {/* Hero */}
      <section className="flex flex-col justify-center items-start w-full h-screen min-h-[680px] px-[5%]">
        <div>
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

      {/* How to Use */}
      <section id="features" className="py-24 px-6">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold mb-4">How to Use</h2>
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
