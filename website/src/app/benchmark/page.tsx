export default function BenchmarkPage() {
  return (
    <main className="min-h-screen py-24 px-6">
      <div className="max-w-5xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">Benchmarks</h1>
        <p className="text-[#737373] mb-12">
          Performance comparison: Traditional API vs Velo Gateway
        </p>

        <div className="space-y-8">
          {/* Throughput */}
          <section className="p-8 rounded-xl bg-[#141414] border border-[#262626]">
            <h2 className="text-2xl font-bold mb-6">Throughput</h2>
            <div className="space-y-4">
              <div>
                <div className="flex justify-between mb-2">
                  <span>Tradicional (Express.js)</span>
                  <span className="text-[#737373]">1,000 req/s</span>
                </div>
                <div className="h-4 bg-[#262626] rounded-full overflow-hidden">
                  <div className="h-full bg-[#737373] rounded-full" style={{ width: "2%" }} />
                </div>
              </div>
              <div>
                <div className="flex justify-between mb-2">
                  <span>Velo Gateway</span>
                  <span className="text-[#00DC82]">50,000 req/s</span>
                </div>
                <div className="h-4 bg-[#262626] rounded-full overflow-hidden">
                  <div className="h-full bg-[#00DC82] rounded-full" style={{ width: "100%" }} />
                </div>
              </div>
            </div>
          </section>

          {/* Latency */}
          <section className="p-8 rounded-xl bg-[#141414] border border-[#262626]">
            <h2 className="text-2xl font-bold mb-6">Latência P99</h2>
            <div className="space-y-4">
              <div>
                <div className="flex justify-between mb-2">
                  <span>Tradicional</span>
                  <span className="text-[#737373]">45ms</span>
                </div>
                <div className="h-4 bg-[#262626] rounded-full overflow-hidden">
                  <div className="h-full bg-[#737373] rounded-full" style={{ width: "100%" }} />
                </div>
              </div>
              <div>
                <div className="flex justify-between mb-2">
                  <span>Velo Gateway</span>
                  <span className="text-[#00DC82]">12ms</span>
                </div>
                <div className="h-4 bg-[#262626] rounded-full overflow-hidden">
                  <div className="h-full bg-[#00DC82] rounded-full" style={{ width: "27%" }} />
                </div>
              </div>
            </div>
          </section>

          {/* Memory */}
          <section className="p-8 rounded-xl bg-[#141414] border border-[#262626]">
            <h2 className="text-2xl font-bold mb-6">Uso de Memória</h2>
            <div className="space-y-4">
              <div>
                <div className="flex justify-between mb-2">
                  <span>Tradicional (Node.js)</span>
                  <span className="text-[#737373]">256 MB</span>
                </div>
                <div className="h-4 bg-[#262626] rounded-full overflow-hidden">
                  <div className="h-full bg-[#737373] rounded-full" style={{ width: "100%" }} />
                </div>
              </div>
              <div>
                <div className="flex justify-between mb-2">
                  <span>Velo Gateway (Go)</span>
                  <span className="text-[#00DC82]">45 MB</span>
                </div>
                <div className="h-4 bg-[#262626] rounded-full overflow-hidden">
                  <div className="h-full bg-[#00DC82] rounded-full" style={{ width: "18%" }} />
                </div>
              </div>
            </div>
          </section>

          {/* Cold Start */}
          <section className="p-8 rounded-xl bg-[#141414] border border-[#262626]">
            <h2 className="text-2xl font-bold mb-6">Cold Start</h2>
            <div className="space-y-4">
              <div>
                <div className="flex justify-between mb-2">
                  <span>Tradicional</span>
                  <span className="text-[#737373]">2.5s</span>
                </div>
                <div className="h-4 bg-[#262626] rounded-full overflow-hidden">
                  <div className="h-full bg-[#737373] rounded-full" style={{ width: "100%" }} />
                </div>
              </div>
              <div>
                <div className="flex justify-between mb-2">
                  <span>Velo Gateway</span>
                  <span className="text-[#00DC82]">15ms</span>
                </div>
                <div className="h-4 bg-[#262626] rounded-full overflow-hidden">
                  <div className="h-full bg-[#00DC82] rounded-full" style={{ width: "1%" }} />
                </div>
              </div>
            </div>
          </section>
        </div>
      </div>
    </main>
  );
}
