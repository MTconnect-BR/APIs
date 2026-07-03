"use client";

export function BenchmarkChart() {
  const metrics = [
    { label: "Latência", traditional: 45, velo: 12, unit: "ms", improvement: "73%" },
    { label: "Throughput", traditional: 1000, velo: 50000, unit: "req/s", improvement: "50x" },
    { label: "Memória", traditional: 256, velo: 45, unit: "MB", improvement: "82%" },
    { label: "CPU", traditional: 80, velo: 15, unit: "%", improvement: "81%" },
  ];

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
      {metrics.map((m, i) => (
        <div key={i} className="p-6 rounded-xl bg-[#0A0A0A] border border-[#262626]">
          <div className="flex justify-between items-center mb-4">
            <h3 className="font-semibold">{m.label}</h3>
            <span className="text-[#00DC82] text-sm font-medium">-{m.improvement}</span>
          </div>
          <div className="space-y-3">
            <div>
              <div className="flex justify-between text-sm mb-1">
                <span className="text-[#737373]">Tradicional</span>
                <span>{m.traditional} {m.unit}</span>
              </div>
              <div className="h-2 bg-[#262626] rounded-full overflow-hidden">
                <div
                  className="h-full bg-[#737373] rounded-full"
                  style={{ width: "100%" }}
                />
              </div>
            </div>
            <div>
              <div className="flex justify-between text-sm mb-1">
                <span className="text-[#737373]">Velo</span>
                <span>{m.velo} {m.unit}</span>
              </div>
              <div className="h-2 bg-[#262626] rounded-full overflow-hidden">
                <div
                  className="h-full bg-[#00DC82] rounded-full"
                  style={{ width: `${(m.velo / m.traditional) * 100}%` }}
                />
              </div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
