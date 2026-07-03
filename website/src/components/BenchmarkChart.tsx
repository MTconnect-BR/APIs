"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Progress } from "@/components/ui/progress";

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
        <Card key={i}>
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-sm font-medium">{m.label}</CardTitle>
            <span className="text-primary text-sm font-medium">-{m.improvement}</span>
          </CardHeader>
          <CardContent className="space-y-4">
            <div>
              <div className="flex justify-between text-sm mb-2">
                <span className="text-muted-foreground">Tradicional</span>
                <span>{m.traditional} {m.unit}</span>
              </div>
              <Progress value={100} className="h-2" />
            </div>
            <div>
              <div className="flex justify-between text-sm mb-2">
                <span className="text-muted-foreground">Velo</span>
                <span>{m.velo} {m.unit}</span>
              </div>
              <Progress value={(m.velo / m.traditional) * 100} className="h-2" />
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
