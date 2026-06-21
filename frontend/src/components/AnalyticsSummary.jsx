export default function AnalyticsSummary({
  analytics,
}) {
  return (
    <div className="border rounded p-4">
      <h2>
        URL Analytics
      </h2>

      <p>
        Short Code:
        {" "}
        {analytics.short_code}
      </p>

      <p>
        Total Clicks:
        {" "}
        {analytics.total_clicks}
      </p>

      <p>
        Unique Visitors:
        {" "}
        {analytics.unique_visitors}
      </p>
    </div>
  );
}