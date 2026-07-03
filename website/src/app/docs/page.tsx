import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export default function DocsPage() {
  return (
    <main className="min-h-screen py-24 px-6">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">Documentação</h1>
        <p className="text-muted-foreground mb-12">
          Comece a usar Velo em minutos.
        </p>

        <div className="space-y-12">
          {/* Installation */}
          <section>
            <h2 className="text-2xl font-bold mb-4">Instalação</h2>
            <Card>
              <CardContent className="pt-6">
                <pre className="text-sm">
{`# Download binário
curl -sSL https://github.com/MTconnect-BR/APIs/releases/latest/download/velo-linux-amd64 -o velo
chmod +x velo

# Ou via Docker
docker pull velo-api/velo:latest

# Ou compile from source
go install github.com/MTconnect-BR/APIs/velo/cmd/velo@latest`}
                </pre>
              </CardContent>
            </Card>
          </section>

          {/* Quick Start */}
          <section>
            <h2 className="text-2xl font-bold mb-4">Quick Start</h2>
            <Card>
              <CardContent className="pt-6">
                <pre className="text-sm">
{`# 1. Create config
cp configs/velo.example.yaml configs/velo.yaml

# 2. Edit config
vim configs/velo.yaml

# 3. Run
./velo configs/velo.yaml`}
                </pre>
              </CardContent>
            </Card>
          </section>

          {/* Configuration */}
          <section>
            <h2 className="text-2xl font-bold mb-4">Configuração</h2>
            <div className="space-y-4">
              <Card>
                <CardHeader>
                  <CardTitle>Server</CardTitle>
                </CardHeader>
                <CardContent>
                  <pre className="text-sm">
{`server:
  host: 0.0.0.0
  port: 8080`}
                  </pre>
                </CardContent>
              </Card>
              <Card>
                <CardHeader>
                  <CardTitle>Rate Limiting</CardTitle>
                </CardHeader>
                <CardContent>
                  <pre className="text-sm">
{`ratelimit:
  enabled: true
  default: "100/min"
  rules:
    - path: "/api/v1/*"
      limit: "1000/min"`}
                  </pre>
                </CardContent>
              </Card>
              <Card>
                <CardHeader>
                  <CardTitle>Cache</CardTitle>
                </CardHeader>
                <CardContent>
                  <pre className="text-sm">
{`cache:
  enabled: true
  driver: redis
  redis:
    addr: localhost:6379
  ttl: 5m`}
                  </pre>
                </CardContent>
              </Card>
              <Card>
                <CardHeader>
                  <CardTitle>Auth</CardTitle>
                </CardHeader>
                <CardContent>
                  <pre className="text-sm">
{`auth:
  enabled: true
  providers:
    - type: jwt
      secret: "\${JWT_SECRET}"
    - type: apikey
      header: "X-API-Key"`}
                  </pre>
                </CardContent>
              </Card>
            </div>
          </section>

          {/* Docker */}
          <section>
            <h2 className="text-2xl font-bold mb-4">Docker</h2>
            <Card>
              <CardContent className="pt-6">
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
              </CardContent>
            </Card>
          </section>
        </div>
      </div>
    </main>
  );
}
