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
    <div className="overflow-x-auto">
      <table className="w-full">
        <thead>
          <tr className="border-b border-[#262626]">
            <th className="text-left py-4 px-6 text-[#737373] font-medium">Feature</th>
            <th className="text-left py-4 px-6 text-[#737373] font-medium">Tradicional</th>
            <th className="text-left py-4 px-6 text-[#00DC82] font-medium">Velo</th>
          </tr>
        </thead>
        <tbody>
          {features.map((row, i) => (
            <tr key={i} className="border-b border-[#262626] hover:bg-[#00DC82]/5">
              <td className="py-4 px-6 font-medium">{row.feature}</td>
              <td className="py-4 px-6 text-[#737373]">
                {typeof row.traditional === "boolean" ? (
                  row.traditional ? "✅" : "❌"
                ) : (
                  row.traditional
                )}
              </td>
              <td className="py-4 px-6 text-[#00DC82]">
                {typeof row.velo === "boolean" ? (
                  row.velo ? "✅" : "❌"
                ) : (
                  row.velo
                )}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
