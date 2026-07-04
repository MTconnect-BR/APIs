"use client";

import { useState, useRef, useCallback } from "react";
import Link from "next/link";

interface TestResult {
  totalRequests: number;
  successCount: number;
  errorCount: number;
  rateLimitedCount: number;
  avgLatency: number;
  p50Latency: number;
  p95Latency: number;
  p99Latency: number;
  throughput: number;
  totalTime: number;
  cacheHits: number;
  bytesTransferred: number;
  latencies: number[];
  timeline: { time: number; requests: number; errors: number }[];
}

interface TestConfig {
  requestCount: number;
  concurrency: number;
  methods: string[];
}

export default function Playground() {
  const [config, setConfig] = useState<TestConfig>({
    requestCount: 50,
    concurrency: 5,
    methods: ["GET"],
  });
  const [isTesting, setIsTesting] = useState<"traditional" | "velo" | null>(null);
  const [progress, setProgress] = useState(0);
  const [traditionalResult, setTraditionalResult] = useState<TestResult | null>(null);
  const [veloResult, setVeloResult] = useState<TestResult | null>(null);
  const [liveMetrics, setLiveMetrics] = useState({
    total: 0,
    success: 0,
    errors: 0,
    rateLimited: 0,
    avgLatency: 0,
  });

  const abortControllerRef = useRef<AbortController | null>(null);

  const TRADITIONAL_API = "https://jsonplaceholder.typicode.com";
  const VELO_API = process.env.NEXT_PUBLIC_VELO_API_URL || "https://velo-api-production.up.railway.app";

  const simulateRequest = async (
    method: string,
    baseUrl: string,
    signal: AbortSignal
  ): Promise<{ success: boolean; latency: number; statusCode: number; bytes: number }> => {
    const start = performance.now();

    let url = baseUrl;
    let body: BodyInit | undefined;

    switch (method) {
      case "GET":
        url += `/posts/${Math.floor(Math.random() * 100) + 1}`;
        break;
      case "POST":
        url += "/posts";
        body = JSON.stringify({
          title: "Test",
          body: "Performance test",
          userId: 1,
        });
        break;
      case "PUT":
        url += `/posts/${Math.floor(Math.random() * 100) + 1}`;
        body = JSON.stringify({
          id: 1,
          title: "Updated",
          body: "Updated content",
          userId: 1,
        });
        break;
      case "DELETE":
        url += `/posts/${Math.floor(Math.random() * 100) + 1}`;
        break;
      case "PATCH":
        url += `/posts/${Math.floor(Math.random() * 100) + 1}`;
        body = JSON.stringify({ title: "Patched" });
        break;
    }

    try {
      const response = await fetch(url, {
        method,
        headers: {
          "Content-Type": "application/json",
        },
        body,
        signal,
      });

      const data = await response.text();
      const latency = performance.now() - start;

      return {
        success: response.ok,
        latency,
        statusCode: response.status,
        bytes: new TextEncoder().encode(data).length,
      };
    } catch (error) {
      const latency = performance.now() - start;
      return {
        success: false,
        latency,
        statusCode: 0,
        bytes: 0,
      };
    }
  };

  const runTest = async (
    type: "traditional" | "velo"
  ): Promise<TestResult> => {
    const baseUrl = type === "traditional" ? TRADITIONAL_API : VELO_API;
    const results: number[] = [];
    const timeline: { time: number; requests: number; errors: number }[] = [];
    let successCount = 0;
    let errorCount = 0;
    let rateLimitedCount = 0;
    let bytesTransferred = 0;

    const controller = new AbortController();
    abortControllerRef.current = controller;

    const startTime = performance.now();
    const batchSize = config.concurrency;
    const totalBatches = Math.ceil(config.requestCount / batchSize);

    for (let batch = 0; batch < totalBatches; batch++) {
      if (controller.signal.aborted) break;

      const batchSizeActual = Math.min(
        batchSize,
        config.requestCount - batch * batchSize
      );

      const promises = Array.from({ length: batchSizeActual }, () => {
        const method = config.methods[Math.floor(Math.random() * config.methods.length)];
        return simulateRequest(method, baseUrl, controller.signal);
      });

      const batchResults = await Promise.all(promises);
      const currentTime = performance.now() - startTime;

      batchResults.forEach((result) => {
        results.push(result.latency);
        bytesTransferred += result.bytes;

        if (result.success) {
          successCount++;
        } else {
          errorCount++;
          if (result.statusCode === 429) {
            rateLimitedCount++;
          }
        }
      });

      const batchSuccess = batchResults.filter((r) => r.success).length;
      const batchErrors = batchResults.filter((r) => !r.success).length;

      timeline.push({
        time: currentTime,
        requests: batchSuccess,
        errors: batchErrors,
      });

      setProgress(((batch + 1) / totalBatches) * 100);
      setLiveMetrics({
        total: (batch + 1) * batchSize,
        success: successCount,
        errors: errorCount,
        rateLimited: rateLimitedCount,
        avgLatency: results.reduce((a, b) => a + b, 0) / results.length,
      });
    }

    const totalTime = performance.now() - startTime;
    const sortedLatencies = [...results].sort((a, b) => a - b);

    return {
      totalRequests: results.length,
      successCount,
      errorCount,
      rateLimitedCount,
      avgLatency: results.reduce((a, b) => a + b, 0) / results.length,
      p50Latency: sortedLatencies[Math.floor(sortedLatencies.length * 0.5)] || 0,
      p95Latency: sortedLatencies[Math.floor(sortedLatencies.length * 0.95)] || 0,
      p99Latency: sortedLatencies[Math.floor(sortedLatencies.length * 0.99)] || 0,
      throughput: (results.length / totalTime) * 1000,
      totalTime,
      cacheHits: 0,
      bytesTransferred,
      latencies: results,
      timeline,
    };
  };

  const handleTestTraditional = async () => {
    setIsTesting("traditional");
    setProgress(0);
    setTraditionalResult(null);
    const result = await runTest("traditional");
    setTraditionalResult(result);
    setIsTesting(null);
  };

  const handleTestVelo = async () => {
    setIsTesting("velo");
    setProgress(0);
    setVeloResult(null);
    const result = await runTest("velo");
    setVeloResult(result);
    setIsTesting(null);
  };

  const handleStop = () => {
    abortControllerRef.current?.abort();
    setIsTesting(null);
  };

  const toggleMethod = (method: string) => {
    setConfig((prev) => ({
      ...prev,
      methods: prev.methods.includes(method)
        ? prev.methods.filter((m) => m !== method)
        : [...prev.methods, method],
    }));
  };

  return (
    <main className="min-h-screen">
      {/* Header */}
      <header className="sticky top-0 z-50 border-b border-border bg-background/80 backdrop-blur-md">
        <div className="max-w-5xl mx-auto px-6 h-14 flex items-center justify-between">
          <Link href="/" className="flex items-center gap-2 font-bold text-lg">
            Velo
          </Link>
          <nav className="flex items-center gap-6 text-sm text-muted-foreground">
            <a href="/#docs" className="hover:text-foreground transition">Docs</a>
            <a href="/#tests" className="hover:text-foreground transition">Tests</a>
            <a href="/#compare" className="hover:text-foreground transition">Compare</a>
            <Link href="/playground" className="text-foreground font-medium">Playground</Link>
            <a href="https://github.com/MTconnect-BR/APIs" target="_blank" rel="noopener noreferrer" className="hover:text-foreground transition">
              GitHub
            </a>
          </nav>
        </div>
      </header>

      <div className="max-w-5xl mx-auto px-6 py-16">
        {/* Title */}
        <div className="text-center mb-12">
          <h1 className="text-4xl md:text-5xl font-bold mb-4">PLAYGROUND</h1>
          <p className="text-muted-foreground text-lg">Teste a API Velo em tempo real</p>
        </div>

        {/* Configuration */}
        <div className="border border-border rounded-lg p-6 mb-8">
          <h2 className="text-xl font-semibold mb-4">Configuração</h2>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
            {/* Request Count */}
            <div>
              <label className="block text-sm font-medium mb-2">
                Requests: {config.requestCount}
              </label>
              <input
                type="range"
                min="10"
                max="200"
                value={config.requestCount}
                onChange={(e) =>
                  setConfig({ ...config, requestCount: parseInt(e.target.value) })
                }
                className="w-full"
              />
              <div className="flex justify-between text-xs text-muted-foreground mt-1">
                <span>10</span>
                <span>200</span>
              </div>
            </div>

            {/* Concurrency */}
            <div>
              <label className="block text-sm font-medium mb-2">
                Concorrentes: {config.concurrency}
              </label>
              <input
                type="range"
                min="1"
                max="50"
                value={config.concurrency}
                onChange={(e) =>
                  setConfig({ ...config, concurrency: parseInt(e.target.value) })
                }
                className="w-full"
              />
              <div className="flex justify-between text-xs text-muted-foreground mt-1">
                <span>1</span>
                <span>50</span>
              </div>
            </div>
          </div>

          {/* HTTP Methods */}
          <div className="mb-6">
            <label className="block text-sm font-medium mb-2">Métodos HTTP</label>
            <div className="flex flex-wrap gap-2">
              {["GET", "POST", "PUT", "DELETE", "PATCH"].map((method) => (
                <button
                  key={method}
                  onClick={() => toggleMethod(method)}
                  className={`px-4 py-2 rounded-lg font-mono text-sm font-medium transition ${
                    config.methods.includes(method)
                      ? "bg-primary text-white"
                      : "bg-muted text-muted-foreground hover:bg-muted/80"
                  }`}
                >
                  {method}
                </button>
              ))}
            </div>
          </div>

          {/* Test Buttons */}
          <div className="flex flex-wrap gap-4">
            <button
              onClick={handleTestTraditional}
              disabled={isTesting !== null || config.methods.length === 0}
              className="px-6 py-3 bg-red-600 hover:bg-red-700 disabled:bg-red-600/50 text-white rounded-lg font-medium transition flex items-center gap-2"
            >
              {isTesting === "traditional" ? (
                <>
                  <span className="animate-spin">⏳</span> Testando...
                </>
              ) : (
                "⚡ Testar Tradicional"
              )}
            </button>

            <button
              onClick={handleTestVelo}
              disabled={isTesting !== null || config.methods.length === 0}
              className="px-6 py-3 bg-primary hover:bg-primary/90 disabled:bg-primary/50 text-white rounded-lg font-medium transition flex items-center gap-2"
            >
              {isTesting === "velo" ? (
                <>
                  <span className="animate-spin">⏳</span> Testando...
                </>
              ) : (
                "🚀 Testar Velo API"
              )}
            </button>

            {isTesting && (
              <button
                onClick={handleStop}
                className="px-6 py-3 bg-muted hover:bg-muted/80 text-foreground rounded-lg font-medium transition"
              >
                ⏹ Parar
              </button>
            )}
          </div>
        </div>

        {/* Live Results */}
        {isTesting && (
          <div className="border border-border rounded-lg p-6 mb-8">
            <h2 className="text-xl font-semibold mb-4">
              RESULTADOS EM TEMPO REAL
              <span className="ml-2 text-sm font-normal text-muted-foreground">
                ({isTesting === "traditional" ? "API Tradicional" : "Velo API"})
              </span>
            </h2>

            <div className="grid grid-cols-2 md:grid-cols-5 gap-4 mb-6">
              <div className="bg-muted rounded-lg p-4 text-center">
                <div className="text-2xl font-bold">{liveMetrics.total}</div>
                <div className="text-sm text-muted-foreground">Total</div>
              </div>
              <div className="bg-muted rounded-lg p-4 text-center">
                <div className="text-2xl font-bold text-green-500">{liveMetrics.success}</div>
                <div className="text-sm text-muted-foreground">Sucesso</div>
              </div>
              <div className="bg-muted rounded-lg p-4 text-center">
                <div className="text-2xl font-bold text-red-500">{liveMetrics.errors}</div>
                <div className="text-sm text-muted-foreground">Erros</div>
              </div>
              <div className="bg-muted rounded-lg p-4 text-center">
                <div className="text-2xl font-bold text-yellow-500">{liveMetrics.rateLimited}</div>
                <div className="text-sm text-muted-foreground">Rate Limited</div>
              </div>
              <div className="bg-muted rounded-lg p-4 text-center">
                <div className="text-2xl font-bold">{Math.round(liveMetrics.avgLatency)}ms</div>
                <div className="text-sm text-muted-foreground">Latência</div>
              </div>
            </div>

            {/* Progress Bar */}
            <div className="w-full bg-muted rounded-full h-3">
              <div
                className="bg-primary h-3 rounded-full transition-all duration-300"
                style={{ width: `${progress}%` }}
              />
            </div>
            <div className="text-center text-sm text-muted-foreground mt-2">
              {Math.round(progress)}% concluído
            </div>
          </div>
        )}

        {/* Results */}
        {(traditionalResult || veloResult) && (
          <div className="border border-border rounded-lg p-6 mb-8">
            <h2 className="text-xl font-semibold mb-4">COMPARAÇÃO</h2>

            <div className="overflow-x-auto">
              <table className="w-full text-sm">
                <thead>
                  <tr className="border-b border-border">
                    <th className="text-left py-3 px-4 font-medium text-muted-foreground">Métrica</th>
                    <th className="text-left py-3 px-4 font-medium text-red-500">Tradicional</th>
                    <th className="text-left py-3 px-4 font-medium text-primary">Velo API</th>
                  </tr>
                </thead>
                <tbody>
                  <tr className="border-b border-border">
                    <td className="py-3 px-4">Taxa de Sucesso</td>
                    <td className="py-3 px-4">
                      {traditionalResult
                        ? `${((traditionalResult.successCount / traditionalResult.totalRequests) * 100).toFixed(1)}%`
                        : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult
                        ? `${((veloResult.successCount / veloResult.totalRequests) * 100).toFixed(1)}%`
                        : "-"}
                    </td>
                  </tr>
                  <tr className="border-b border-border">
                    <td className="py-3 px-4">Latência Média</td>
                    <td className="py-3 px-4">
                      {traditionalResult ? `${Math.round(traditionalResult.avgLatency)}ms` : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult ? `${Math.round(veloResult.avgLatency)}ms` : "-"}
                    </td>
                  </tr>
                  <tr className="border-b border-border">
                    <td className="py-3 px-4">P50 Latência</td>
                    <td className="py-3 px-4">
                      {traditionalResult ? `${Math.round(traditionalResult.p50Latency)}ms` : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult ? `${Math.round(veloResult.p50Latency)}ms` : "-"}
                    </td>
                  </tr>
                  <tr className="border-b border-border">
                    <td className="py-3 px-4">P95 Latência</td>
                    <td className="py-3 px-4">
                      {traditionalResult ? `${Math.round(traditionalResult.p95Latency)}ms` : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult ? `${Math.round(veloResult.p95Latency)}ms` : "-"}
                    </td>
                  </tr>
                  <tr className="border-b border-border">
                    <td className="py-3 px-4">P99 Latência</td>
                    <td className="py-3 px-4">
                      {traditionalResult ? `${Math.round(traditionalResult.p99Latency)}ms` : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult ? `${Math.round(veloResult.p99Latency)}ms` : "-"}
                    </td>
                  </tr>
                  <tr className="border-b border-border">
                    <td className="py-3 px-4">Throughput</td>
                    <td className="py-3 px-4">
                      {traditionalResult ? `${Math.round(traditionalResult.throughput)} req/s` : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult ? `${Math.round(veloResult.throughput)} req/s` : "-"}
                    </td>
                  </tr>
                  <tr className="border-b border-border">
                    <td className="py-3 px-4">Rate Limited</td>
                    <td className="py-3 px-4">
                      {traditionalResult ? traditionalResult.rateLimitedCount : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult ? veloResult.rateLimitedCount : "-"}
                    </td>
                  </tr>
                  <tr className="border-b border-border">
                    <td className="py-3 px-4">Tempo Total</td>
                    <td className="py-3 px-4">
                      {traditionalResult ? `${(traditionalResult.totalTime / 1000).toFixed(2)}s` : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult ? `${(veloResult.totalTime / 1000).toFixed(2)}s` : "-"}
                    </td>
                  </tr>
                  <tr>
                    <td className="py-3 px-4">Bytes Transferidos</td>
                    <td className="py-3 px-4">
                      {traditionalResult
                        ? `${(traditionalResult.bytesTransferred / 1024).toFixed(1)} KB`
                        : "-"}
                    </td>
                    <td className="py-3 px-4">
                      {veloResult
                        ? `${(veloResult.bytesTransferred / 1024).toFixed(1)} KB`
                        : "-"}
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        )}

        {/* Info */}
        <div className="border border-border rounded-lg p-6">
          <h2 className="text-xl font-semibold mb-4">Como funciona</h2>
          <div className="space-y-4 text-muted-foreground">
            <p>
              <strong className="text-foreground">API Tradicional:</strong> Requests diretos para
              APIs públicas (jsonplaceholder.typicode.com) sem otimizações de gateway.
            </p>
            <p>
              <strong className="text-foreground">Velo API:</strong> Requests passando pelo
              gateway Velo com rate limiting otimizado, cache de respostas e balanceamento de carga.
            </p>
            <p>
              <strong className="text-foreground">Métricas:</strong> Todas as métricas são coletadas
              em tempo real via <code className="bg-muted px-1 rounded">performance.now()</code> e
              headers HTTP.
            </p>
          </div>
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
