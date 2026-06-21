export default function AnalyticsSummary({
  analytics,
}) {
  const details = [
    {
      label: "Short Code",
      value: analytics.short_code,
    },
    {
      label: "Total Clicks",
      value: analytics.total_clicks,
    },
    {
      label: "Unique Visitors",
      value: analytics.unique_visitors,
    },
    {
      label: "Tracking Status",
      value: "Active",
    },
  ];

  return (
    <div className="url-details-card">

      <div className="url-details-header">

        <h2>
          URL Details
        </h2>

        <p>
          Metadata for the selected
          shortened URL.
        </p>

      </div>

      <div className="details-list">

        {details.map((item) => (
          <div
            key={item.label}
            className="details-row"
          >

            <span className="details-label">
              {item.label}
            </span>

            <span className="details-value">
              {item.value}
            </span>

          </div>
        ))}

      </div>

    </div>
  );
}