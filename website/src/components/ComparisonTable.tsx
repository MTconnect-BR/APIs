import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

export function ComparisonTable() {
  const features = [
    { feature: "Rate Limiting", traditional: false, velo: true },
    { feature: "Cache Global", traditional: false, velo: true },
    { feature: "Autenticação", traditional: "Manual", velo: "JWT + OAuth2" },
    { feature: "Load Balancing", traditional: false, velo: true },
    { feature: "Observabilidade", traditional: "Logs básicos", velo: "Prometheus + Tracing" },
    { feature: "Documentação", traditional: "Manual", velo: "Auto OpenAPI 3.1" },
    { feature: "Versionamento", traditional: "URL path", velo: "Header + Path + Query" },
    { feature: "Deploy", traditional: "Complexo", velo: "Binário único" },
    { feature: "Performance", traditional: "~50ms", velo: "~12ms" },
    { feature: "Throughput", traditional: "1k req/s", velo: "50k req/s" },
  ];

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>Feature</TableHead>
          <TableHead>Tradicional</TableHead>
          <TableHead className="text-primary">Velo</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {features.map((row, i) => (
          <TableRow key={i}>
            <TableCell className="font-medium">{row.feature}</TableCell>
            <TableCell className="text-muted-foreground">
              {typeof row.traditional === "boolean" ? (
                row.traditional ? "✅" : "❌"
              ) : (
                row.traditional
              )}
            </TableCell>
            <TableCell className="text-primary">
              {typeof row.velo === "boolean" ? (
                row.velo ? "✅" : "❌"
              ) : (
                row.velo
              )}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}
