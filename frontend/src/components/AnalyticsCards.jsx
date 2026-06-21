export default function AnalyticsCards({
  totalClicks,
  uniqueVisitors,
}) {
  const cards = [
    {
      title: "Total Clicks",
      value: totalClicks,
      description:
        "All tracked visits",
    },
    {
      title: "Unique Visitors",
      value: uniqueVisitors,
      description:
        "Distinct visitors",
    },
  ];

  return (
    <div className="analytics-cards-grid">
      {cards.map((card) => (
        <div
          key={card.title}
          className="analytics-kpi-card"
        >
          <span className="kpi-label">
            {card.title}
          </span>

          <h2 className="kpi-value">
            {card.value}
          </h2>

          <p className="kpi-description">
            {card.description}
          </p>
        </div>
      ))}
    </div>
  );
}