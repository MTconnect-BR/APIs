export default function DocsPage() {
  return (
    <main className="min-h-screen py-24 px-6">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">Documentação</h1>
        <p className="text-[#737373] mb-12">
          Comece a usar Velo em minutos.
        </p>

        <div className="space-y-12">
          {/* Installation */}
          <section>
            <h2 className="text-2xl font-bold mb-4">Instalação</h2>
            <div className="p-4 rounded-lg bg-[#141414] border border-[#262626]">
              <pre className="text-sm">
{`# Download binário
curl -sSL https://github.com/velo-api/velo/releases/latest/download/velo-linux-amd64 -o velo
chmod +x velo

# Ou via Docker
docker pull velo-api/velo:latest

# Ou compile from source
go install github.com/velo-api/velo/cmd/velo@latest`}
              </pre>
            </div>
          </section>

          {/* Quick Start */}
          <section>
            <h2 className="text-2xl font-bold mb-4">Quick Start</h2>
            <div className="p-4 rounded-lg bg-[#141414] border border-[#262626]">
              <pre className="text-sm">
{`# 1. Create config
cp configs/velo.example.yaml configs/velo.yaml

# 2. Edit config
vim configs/velo.yaml

# 3. Run
./velo configs/velo.yaml`}
              </pre>
            </div>
          </section>

          {/* Configuration */}
          <section>
            <h2 className="text-2xl font-bold mb-4">Configuração</h2>
            <div className="space-y-4">
              <div className="p-4 rounded-lg bg-[#141414] border border-[#262626]">
                <h3 className="font-semibold mb-2">Server</h3>
                <pre className="text-sm">
{`server:
  host: 0.0.0.0
  port: 8080`}
                </pre>
              </div>
              <div className="p-4 rounded-lg bg-[#141414] border border-[#262626]">
                <h3 className="font-semibold mb-2">Rate Limiting</h3>
                <pre className="text-sm">
{`ratelimit:
  enabled: true
  default: "100/min"
  rules:
    - path: "/api/v1/*"
      limit: "1000/min"`}
                </pre>
              </div>
              <div className="p-4 rounded-lg bg-[#141414] border border-[#262626]">
                <h3 className="font-semibold mb-2">Cache</h3>
                <pre className="text-sm">
{`cache:
  enabled: true
  driver: redis
  redis:
    addr: localhost:6379
  ttl: 5m`}
                </pre>
              </div>
              <div className="p-4 rounded-lg bg-[#141414] border border-[#262626]">
                <h3 className="font-semibold mb-2">Auth</h3>
                <pre className="text-sm">
{`auth:
  enabled: true
  providers:
    - type: jwt
      secret: "\${JWT_SECRET}"
    - type: apikey
      header: "X-API-Key"`}
                </pre>
              </div>
            </div>
          </section>

          {/* Docker */}
          <section>
            <h2 className="text-2xl font-bold mb-4">Docker</h2>
            <div className="p-4 rounded-lg bg-[#141414] border border-[#262626]">
              <pre className="text-sm">
{`# Run with Docker
docker run -p 8080:8080 -v ./configs:/app/configs velo-api/velo

# Docker Compose
version: '3.8'
services:
  velo:
    image: velo-api/velo
    ports:
      - "8080:8080"
    volumes:
      - ./configs:/app/configs
  redis:
    image: redis:alpine
    ports:
      - "6379:6379"`}
              </pre>
            </div>
          </section>
        </div>
      </div>
    </main>
  );
}
