import Link from "next/link";

export default function Docs() {
  return (
    <main className="min-h-screen">
      {/* Header */}
      <header className="sticky top-0 z-50 border-b border-border bg-background/80 backdrop-blur-md">
        <div className="max-w-5xl mx-auto px-6 h-14 flex items-center justify-between">
          <Link href="/" className="flex items-center gap-2 font-bold text-lg">
            Velo
          </Link>
          <nav className="flex items-center gap-6 text-sm text-muted-foreground">
            <Link href="/#docs" className="hover:text-foreground transition">Docs</Link>
            <a href="/#tests" className="hover:text-foreground transition">Tests</a>
            <a href="/#compare" className="hover:text-foreground transition">Compare</a>
            <Link href="/playground" className="hover:text-foreground transition">Playground</Link>
            <a href="https://github.com/MTconnect-BR/APIs" target="_blank" rel="noopener noreferrer" className="hover:text-foreground transition">
              GitHub
            </a>
          </nav>
        </div>
      </header>

      <div className="max-w-5xl mx-auto px-6 py-16">
        {/* Title */}
        <div className="text-center mb-12">
          <h1 className="text-4xl md:text-5xl font-bold mb-4">GUIA COMPLETO</h1>
          <p className="text-muted-foreground text-lg">Aprenda sobre APIs de forma simples</p>
        </div>

        {/* O que é uma API */}
        <section className="mb-16">
          <h2 className="text-2xl font-bold mb-6">O que é uma API?</h2>
          <div className="bg-[#272931] rounded-lg p-8 text-lg leading-relaxed">
            <p className="mb-4">
              Imagine que você está num restaurante. Você senta à mesa, olha o cardápio e escolhe o que quer comer.
              Um <strong className="text-foreground">garçom</strong> anota seu pedido e leva para a cozinha.
            </p>
            <p className="mb-4">
              A <strong className="text-primary">API</strong> é como esse <strong className="text-foreground">garçom</strong>:
            </p>
            <ul className="list-disc list-inside space-y-2 text-muted-foreground ml-4">
              <li>Você (o <strong className="text-foreground">cliente</strong>) faz um pedido</li>
              <li>O garçom (<strong className="text-primary">API</strong>) leva o pedido para a cozinha</li>
              <li>A cozinha (<strong className="text-foreground">servidor</strong>) prepara o pedido</li>
              <li>O garçom traz de volta o que você pediu</li>
            </ul>
            <p className="mt-4">
              <strong className="text-foreground">API</strong> significa <strong className="text-primary">Application Programming Interface</strong> —
              ou seja, é a forma como dois programas se comunicam pela internet.
            </p>
          </div>
        </section>

        {/* Como funciona */}
        <section className="mb-16">
          <h2 className="text-2xl font-bold mb-6">Como funciona uma API?</h2>
          <div className="space-y-6">
            <div className="border border-border rounded-lg p-6">
              <h3 className="font-semibold mb-3 text-primary">1. Request (Pedido)</h3>
              <p className="text-muted-foreground mb-3">
                Quando você abre um site, seu navegador faz um <strong className="text-foreground">pedido</strong> para
                o servidor. É como dizer: <code className="bg-[#272931] px-1 rounded">"Me dê os dados do usuário 1"</code>
              </p>
              <div className="bg-[#272931] rounded-lg p-4 font-mono text-sm">
                <pre>{`GET https://api.example.com/users/1

Headers:
  Content-Type: application/json
  Authorization: Bearer sk-abc123`}</pre>
              </div>
            </div>

            <div className="border border-border rounded-lg p-6">
              <h3 className="font-semibold mb-3 text-primary">2. Response (Resposta)</h3>
              <p className="text-muted-foreground mb-3">
                O servidor processa o pedido e devolve uma <strong className="text-foreground">resposta</strong> com
                os dados pedidos. É como o garçom trazer o prato que você pediu!
              </p>
              <div className="bg-[#272931] rounded-lg p-4 font-mono text-sm">
                <pre>{`HTTP/1.1 200 OK
Content-Type: application/json

{
  "id": 1,
  "name": "Maria",
  "email": "maria@example.com"
}`}</pre>
              </div>
            </div>

            <div className="border border-border rounded-lg p-6">
              <h3 className="font-semibold mb-3 text-primary">3. Métodos (Tipos de Pedido)</h3>
              <p className="text-muted-foreground mb-3">
                Existem diferentes tipos de pedidos que você pode fazer:
              </p>
              <div className="grid grid-cols-2 md:grid-cols-3 gap-3">
                <div className="bg-[#272931] rounded-lg p-3 text-center">
                  <div className="text-green-500 font-bold">GET</div>
                  <div className="text-xs text-muted-foreground">Ler dados</div>
                </div>
                <div className="bg-[#272931] rounded-lg p-3 text-center">
                  <div className="text-blue-500 font-bold">POST</div>
                  <div className="text-xs text-muted-foreground">Criar dados</div>
                </div>
                <div className="bg-[#272931] rounded-lg p-3 text-center">
                  <div className="text-yellow-500 font-bold">PUT</div>
                  <div className="text-xs text-muted-foreground">Atualizar tudo</div>
                </div>
                <div className="bg-[#272931] rounded-lg p-3 text-center">
                  <div className="text-orange-500 font-bold">PATCH</div>
                  <div className="text-xs text-muted-foreground">Atualizar parcial</div>
                </div>
                <div className="bg-[#272931] rounded-lg p-3 text-center">
                  <div className="text-red-500 font-bold">DELETE</div>
                  <div className="text-xs text-muted-foreground">Deletar dados</div>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* Status Codes */}
        <section className="mb-16">
          <h2 className="text-2xl font-bold mb-6">Códigos de Status</h2>
          <p className="text-muted-foreground mb-6">
            Toda resposta vem com um <strong className="text-foreground">código de status</strong> que diz se deu certo ou não.
            É como o garçom dizer se seu pedido foi aceito ou não!
          </p>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="border border-green-500/30 rounded-lg p-4">
              <div className="text-green-500 font-bold text-lg mb-2">2xx — Sucesso</div>
              <div className="bg-[#272931] rounded-lg p-3 font-mono text-sm space-y-1">
                <div><span className="text-green-500">200</span> OK — Tudo certo!</div>
                <div><span className="text-green-500">201</span> Created — Criado!</div>
                <div><span className="text-green-500">204</span> No Content — Sem conteúdo</div>
              </div>
            </div>
            <div className="border border-yellow-500/30 rounded-lg p-4">
              <div className="text-yellow-500 font-bold text-lg mb-2">3xx — Redirecionamento</div>
              <div className="bg-[#272931] rounded-lg p-3 font-mono text-sm space-y-1">
                <div><span className="text-yellow-500">301</span> Moved — Mudou de lugar</div>
                <div><span className="text-yellow-500">304</span> Not Modified — Não mudou</div>
              </div>
            </div>
            <div className="border border-orange-500/30 rounded-lg p-4">
              <div className="text-orange-500 font-bold text-lg mb-2">4xx — Erro do Cliente</div>
              <div className="bg-[#272931] rounded-lg p-3 font-mono text-sm space-y-1">
                <div><span className="text-orange-500">400</span> Bad Request — Pedido errado</div>
                <div><span className="text-orange-500">401</span> Unauthorized — Não autenticado</div>
                <div><span className="text-orange-500">403</span> Forbidden — Proibido</div>
                <div><span className="text-orange-500">404</span> Not Found — Não encontrado</div>
                <div><span className="text-orange-500">429</span> Too Many Requests — Muitos pedidos!</div>
              </div>
            </div>
            <div className="border border-red-500/30 rounded-lg p-4">
              <div className="text-red-500 font-bold text-lg mb-2">5xx — Erro do Servidor</div>
              <div className="bg-[#272931] rounded-lg p-3 font-mono text-sm space-y-1">
                <div><span className="text-red-500">500</span> Internal Error — Erro interno</div>
                <div><span className="text-red-500">502</span> Bad Gateway — Gateway ruim</div>
                <div><span className="text-red-500">503</span> Unavailable — Indisponível</div>
              </div>
            </div>
          </div>
        </section>

        {/* O que é um API Gateway */}
        <section className="mb-16">
          <h2 className="text-2xl font-bold mb-6">O que é um API Gateway?</h2>
          <div className="bg-[#272931] rounded-lg p-8">
            <p className="text-lg leading-relaxed mb-4">
              Um <strong className="text-primary">API Gateway</strong> é como um <strong className="text-foreground">porteiro</strong> do
              prédio. Todos os pedidos passam por ele antes de chegar ao servidor.
            </p>
            <p className="text-muted-foreground mb-6">O que o gateway faz:</p>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="bg-background rounded-lg p-4 border border-border">
                <div className="font-semibold mb-2">🛡️ Rate Limiting</div>
                <p className="text-sm text-muted-foreground">Limita quantos pedidos alguém pode fazer por minuto. É como dizer "só 10 pedidos por pessoa!"</p>
              </div>
              <div className="bg-background rounded-lg p-4 border border-border">
                <div className="font-semibold mb-2">⚡ Cache</div>
                <p className="text-sm text-muted-foreground">Guarda respostas frequentes para não precisar perguntar de novo. Como lembrar do pedido do cliente!</p>
              </div>
              <div className="bg-background rounded-lg p-4 border border-border">
                <div className="font-semibold mb-2">🔑 Auth</div>
                <p className="text-sm text-muted-foreground">Verifica se você tem permissão. É como pedir seu RG na entrada!</p>
              </div>
              <div className="bg-background rounded-lg p-4 border border-border">
                <div className="font-semibold mb-2">⚖️ Load Balancing</div>
                <p className="text-sm text-muted-foreground">Distribui os pedidos entre vários servidores. Como vários garçons dividindo o trabalho!</p>
              </div>
            </div>
          </div>
        </section>

        {/* Por que usar o Velo */}
        <section className="mb-16">
          <h2 className="text-2xl font-bold mb-6">Por que usar o Velo?</h2>
          <div className="space-y-4">
            <div className="flex items-start gap-4 border border-border rounded-lg p-6">
              <div className="text-2xl">🚀</div>
              <div>
                <h3 className="font-semibold mb-1">Rápido como um foguete</h3>
                <p className="text-muted-foreground">Feito em Go, processa milhares de requests por segundo com latência mínima.</p>
              </div>
            </div>
            <div className="flex items-start gap-4 border border-border rounded-lg p-6">
              <div className="text-2xl">🆓</div>
              <div>
                <h3 className="font-semibold mb-1">100% Gratuito e Open Source</h3>
                <p className="text-muted-foreground">Sem taxas escondidas. Código aberto para você ver e modificar como quiser.</p>
              </div>
            </div>
            <div className="flex items-start gap-4 border border-border rounded-lg p-6">
              <div className="text-2xl">🔧</div>
              <div>
                <h3 className="font-semibold mb-1">Fácil de configurar</h3>
                <p className="text-muted-foreground">Um único arquivo YAML. Copie, cole, rode. Pronto!</p>
              </div>
            </div>
            <div className="flex items-start gap-4 border border-border rounded-lg p-6">
              <div className="text-2xl">📊</div>
              <div>
                <h3 className="font-semibold mb-1">Observabilidade completa</h3>
                <p className="text-muted-foreground">Métricas, logs e traces em tempo real. Saiba exatamente o que está acontecendo.</p>
              </div>
            </div>
          </div>
        </section>

        {/* Quick Start */}
        <section className="mb-16">
          <h2 className="text-2xl font-bold mb-6">Início Rápido</h2>
          <div className="bg-[#272931] rounded-lg p-6 font-mono text-sm space-y-4">
            <div>
              <div className="text-muted-foreground mb-2"># 1. Baixe o binário</div>
              <div className="text-foreground">curl -sSL https://github.com/MTconnect-BR/APIs/releases/latest/download/velo-linux-amd64 -o velo</div>
            </div>
            <div>
              <div className="text-muted-foreground mb-2"># 2. Torne executável</div>
              <div className="text-foreground">chmod +x velo</div>
            </div>
            <div>
              <div className="text-muted-foreground mb-2"># 3. Crie o arquivo de configuração</div>
              <div className="text-foreground">cat &gt; velo.yaml &lt;&lt; EOF</div>
              <pre className="text-primary">{`server:
  port: 8080

ratelimit:
  enabled: true
  default: "100/min"

cache:
  enabled: true
  ttl: 5m

auth:
  enabled: false`}</pre>
              <div className="text-foreground">EOF</div>
            </div>
            <div>
              <div className="text-muted-foreground mb-2"># 4. Rode!</div>
              <div className="text-foreground">./velo</div>
            </div>
          </div>
        </section>

        {/* Back to home */}
        <div className="text-center">
          <Link
            href="/playground"
            className="inline-block bg-primary text-white px-8 py-3 rounded-lg font-medium hover:bg-primary/90 transition"
          >
            Testar no Playground →
          </Link>
        </div>
      </div>

      {/* Footer */}
      <footer className="py-8 px-6 border-t border-border mt-16">
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
