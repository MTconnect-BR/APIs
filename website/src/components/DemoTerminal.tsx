interface DemoTerminalProps {
  title: string;
  description: string;
  code: string;
}

export function DemoTerminal({ title, description, code }: DemoTerminalProps) {
  return (
    <div className="rounded-xl overflow-hidden border border-[#262626]">
      <div className="flex items-center gap-2 px-4 py-3 bg-[#141414] border-b border-[#262626]">
        <div className="flex gap-2">
          <div className="w-3 h-3 rounded-full bg-[#FF5F56]" />
          <div className="w-3 h-3 rounded-full bg-[#FFBD2E]" />
          <div className="w-3 h-3 rounded-full bg-[#27C93F]" />
        </div>
        <span className="text-sm text-[#737373] ml-2">{title}</span>
      </div>
      <div className="p-4 bg-[#0A0A0A]">
        <p className="text-sm text-[#737373] mb-4">{description}</p>
        <pre className="text-sm overflow-x-auto text-[#FAFAFA]">
          {code}
        </pre>
      </div>
    </div>
  );
}
