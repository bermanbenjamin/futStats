import { ChartConfig, ChartContainer } from "@/components/ui/chart";
import { Player } from "@/http/types";
import { Cell, Pie, PieChart } from "recharts";

function StatsChart({ player }: { player: Player }) {
  const chartConfig = {
    desktop: {
      label: "Desktop",
      color: "#2563eb",
    },
    mobile: {
      label: "Mobile",
      color: "#60a5fa",
    },
  } satisfies ChartConfig;

  const data = [
    { label: "Gols", value: player.goals, color: "#2563eb" },
    { label: "Assistências", value: player.assists, color: "#82ca9d" },
    { label: "Dribles", value: player.dribbles, color: "#ef4444" },
    { label: "Desarmes", value: player.disarms, color: "#8884d8" },
  ];

  return (
    <div className="w-full h-full pt-4">
      <h2 className="text-2xl font-bold">Análise de desempenho</h2>
      <ChartContainer config={chartConfig}>
        <PieChart width={800} height={400}>
          {data.every((item) => item.value === 0) ? (
            <text
              x="50%"
              y="50%"
              textAnchor="middle"
              dominantBaseline="middle"
              className="text-sm text-muted-foreground"
            >
              No data available
            </text>
          ) : (
            <Pie
              data={data}
              dataKey="value"
              cx="50%"
              cy="50%"
              outerRadius="80%"
              labelLine={false}
              label={({ label }) => `${label}`}
              nameKey="label"
              valueKey="value"
            >
              {data.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={entry.color} />
              ))}
            </Pie>
          )}
        </PieChart>
      </ChartContainer>
    </div>
  );
}

export default StatsChart;
