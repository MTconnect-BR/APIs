import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ComparisonTable } from "@/components/ComparisonTable";
import { FeatureCard } from "@/components/FeatureCard";
import { BenchmarkChart } from "@/components/BenchmarkChart";

export default function Home() {
  return (
    <main className="min-h-screen">
      {/* Hero */}
      <section className="relative py-32 px-6 overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-b from-primary/5 to-transparent" />
        <div className="relative max-w-5xl mx-auto text-center">
          <Badge variant="outline" className="mb-6">Open Source • Free</Badge>
          <h1 className="text-5xl md:text-7xl font-bold mb-6 tracking-tight">
            APIs lentas estão{" "}
            <span className="text-primary glow-text">custando dinheiro</span>
          </h1>
          <p className="text-xl text-muted-foreground mb-10 max-w-2xl mx-auto">
            Velo resolve os 7 problemas que toda API enfrenta: rate limiting,
            cache, autenticação, load balancing, observabilidade, documentação e
            versionamento.
          </p>
          <div className="flex gap-4 justify-center">
            <Link href="/docs">
              <Button size="lg">Comece agora</Button>
            </Link>
            <Link href="/demo">
              <Button size="lg" variant="outline">Ver demo</Button>
            </Link>
          </div>
        </div>
      </section>

      {/* Problems */}
      <section className="py-24 px-6">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold text-center mb-16">
            7 problemas que toda API enfrenta
          </h2>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <FeatureCard
              icon="⚡"
              title="Rate Limiting"
              description="Controle requisições por IP, chave ou tenant. Token bucket + sliding window."
            />
            <FeatureCard
              icon="📦"
              title="Cache Global"
              description="Redis-backed com TTL, invalidação por tag. Cache hit/miss metrics."
            />
            <FeatureCard
              icon="🔐"
              title="Autenticação"
              description="JWT, OAuth2, API keys, RBAC. Middleware chain flexível."
            />
            <FeatureCard
              icon="⚖️"
              title="Load Balancing"
              description="Round-robin, least connections, weighted. Health checks automáticos."
            />
            <FeatureCard
              icon="📊"
              title="Observabilidade"
              description="Prometheus metrics, structured logs, distributed tracing."
            />
            <FeatureCard
              icon="📝"
              title="Documentação"
              description="OpenAPI 3.1 auto-gerada a partir do config."
            />
            <FeatureCard
              icon="🏷️"
              title="Versionamento"
              description="Header, path ou query-based. Múltiplas versões simultâneas."
            />
          </div>
        </div>
      </section>

      {/* Comparison */}
      <section className="py-24 px-6 bg-card">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold text-center mb-16">
            Tradicional vs Velo
          </h2>
          <ComparisonTable />
        </div>
      </section>

      {/* Benchmark */}
      <section className="py-24 px-6">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl font-bold text-center mb-16">
            Performance em números
          </h2>
          <BenchmarkChart />
        </div>
      </section>

      {/* CTA */}
      <section className="py-24 px-6 bg-card">
        <div className="max-w-3xl mx-auto text-center">
          <h2 className="text-4xl font-bold mb-6">
            Pronto para acelerar?
          </h2>
          <p className="text-muted-foreground mb-10">
            Deploy em minutos. Binário único. Zero dependências.
          </p>
          <Link href="/docs">
            <Button size="lg">Ler documentação</Button>
          </Link>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-12 px-6 border-t border-border">
        <div className="max-w-5xl mx-auto flex justify-between items-center">
          <span className="text-muted-foreground">© 2024 Velo API</span>
          <div className="flex gap-6">
            <a href="https://github.com/MTconnect-BR/APIs" className="text-muted-foreground hover:text-primary transition">
              GitHub
            </a>
            <Link href="/docs" className="text-muted-foreground hover:text-primary transition">
              Docs
            </Link>
          </div>
        </div>
      </footer>
    </main>
  );
}
