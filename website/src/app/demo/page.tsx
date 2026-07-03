import { DemoTerminal } from "@/components/DemoTerminal";

export default function DemoPage() {
  return (
    <main className="min-h-screen py-24 px-6">
      <div className="max-w-5xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">Demo interativo</h1>
        <p className="text-[#737373] mb-12">
          Veja a diferença entre uma API tradicional e uma API com Velo.
        </p>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <DemoTerminal
            title="API Tradicional"
            description="Sem rate limiting, cache, ou auth"
            code={`// Express.js tradicional
app.get('/api/users', (req, res) => {
  const users = await db.query('SELECT * FROM users');
  res.json(users);
});

// Problemas:
// ❌ Sem rate limiting
// ❌ Sem cache
// ❌ Sem auth
// ❌ Sem observabilidade`}
          />
          <DemoTerminal
            title="Com Velo"
            description="Gateway configurado em YAML"
            code={`# velo.yaml
server:
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
      secret: "\${JWT_SECRET}"

# Resultado:
# ✅ Rate limiting automático
# ✅ Cache Redis
# ✅ JWT auth
# ✅ Prometheus metrics`}
          />
        </div>

        <div className="mt-16">
          <h2 className="text-2xl font-bold mb-6">Requisição vs Resposta</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="p-6 rounded-xl bg-[#141414] border border-[#262626]">
              <h3 className="font-semibold mb-4 text-[#737373]">Request</h3>
              <pre className="text-sm overflow-x-auto">
{`GET /api/users HTTP/1.1
Host: api.example.com
Authorization: Bearer eyJhbG...
X-API-Version: v1
X-API-Key: velo_xxx`}
              </pre>
            </div>
            <div className="p-6 rounded-xl bg-[#141414] border border-[#262626]">
              <h3 className="font-semibold mb-4 text-[#00DC82]">Response</h3>
              <pre className="text-sm overflow-x-auto">
{`HTTP/1.1 200 OK
Content-Type: application/json
X-Cache: HIT
X-RateLimit-Remaining: 99
X-API-Version: v1

{
  "data": [...],
  "meta": {
    "cached": true,
    "latency": "12ms"
  }
}`}
              </pre>
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}
