import {
  ResponsiveContainer,
  LineChart,
  Line,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
} from "recharts";

export default function DailyClicksChart({
  data,
}) {

  const formattedData =
    data.map((item) => ({
      ...item,
      date: new Date(
        item.date
      ).toLocaleDateString(
        "en-US",
        {
          month: "short",
          day: "numeric",
        }
      ),
    }));

  return (
    <div
      className="chart-container"
    >
      <ResponsiveContainer
        width="100%"
        height={320}
      >
        <LineChart
          data={formattedData}
        >
          <CartesianGrid
            strokeDasharray="3 3"
          />

          <XAxis
            dataKey="date"
          />

          <YAxis
            allowDecimals={false}
          />

          <Tooltip />

          <Line
            type="monotone"
            dataKey="clicks"
            stroke="#2563eb"
            strokeWidth={3}
            dot={{
              r: 5,
            }}
          />

        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}