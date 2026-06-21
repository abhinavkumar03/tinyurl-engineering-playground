export default function DailyClicksChart({
  data,
}) {
  return (
    <div>
      <h3>Daily Click Trend</h3>

      <pre>
        {JSON.stringify(data, null, 2)}
      </pre>
    </div>
  );
}