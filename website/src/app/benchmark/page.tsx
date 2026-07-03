import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";

export default function BenchmarkPage() {
  const benchmarks = [
    {
      title: "Throughput",
      traditional: { value: 1000, label: "1,000 req/s", percent: 2 },
      velo: { value: 50000, label: "50,000 req/s", percent: 100 },
    },
    {
      title: "Latência P99",
      traditional: { value: 45, label: "45ms", percent: 100 },
      velo: { value: 12, label: "12ms", percent: 27 },
    },
    {
      title: "Uso de Memória",
      traditional: { value: 256, label: "256 MB", percent: 100 },
      velo: { value: 45, label: "45 MB", percent: 18 },
    },
    {
      title: "Cold Start",
      traditional: { value: 2.5, label: "2.5s", percent: 100 },
      velo: { value: 0.015, label: "15ms", percent: 1 },
    },
  ];

  return (
    <main className="min-h-screen py-24 px-6">
      <div className="max-w-5xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">Benchmarks</h1>
        <p className="text-muted-foreground mb-12">
          Performance comparison: Traditional API vs Velo Gateway
        </p>

        <div className="space-y-8">
          {benchmarks.map((b, i) => (
            <Card key={i}>
              <CardHeader>
                <CardTitle>{b.title}</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div>
                  <div className="flex justify-between mb-2">
                    <span>Tradicional</span>
                    <span className="text-muted-foreground">{b.traditional.label}</span>
                  </div>
                  <Progress value={b.traditional.percent} className="h-4" />
                </div>
                <div>
                  <div className="flex justify-between mb-2">
                    <span>Velo Gateway</span>
                    <span className="text-primary">{b.velo.label}</span>
                  </div>
                  <Progress value={b.velo.percent} className="h-4" />
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    </main>
  );
}
