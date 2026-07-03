import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

interface DemoTerminalProps {
  title: string;
  description: string;
  code: string;
}

export function DemoTerminal({ title, description, code }: DemoTerminalProps) {
  return (
    <Card>
      <CardHeader>
        <div className="flex items-center gap-2">
          <div className="flex gap-2">
            <div className="w-3 h-3 rounded-full bg-red-500" />
            <div className="w-3 h-3 rounded-full bg-yellow-500" />
            <div className="w-3 h-3 rounded-full bg-green-500" />
          </div>
          <CardTitle className="text-sm font-medium ml-2">{title}</CardTitle>
        </div>
      </CardHeader>
      <CardContent>
        <p className="text-sm text-muted-foreground mb-4">{description}</p>
        <pre className="text-sm overflow-x-auto">
          {code}
        </pre>
      </CardContent>
    </Card>
  );
}
