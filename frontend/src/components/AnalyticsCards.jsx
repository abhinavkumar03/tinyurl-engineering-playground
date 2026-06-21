export default function AnalyticsCards({
  totalClicks,
  uniqueVisitors,
}) {
  return (
    <div className="grid grid-cols-2 gap-4">
      <div className="border rounded p-4">
        <h3>Total Clicks</h3>

        <p className="text-3xl font-bold">
          {totalClicks}
        </p>
      </div>

      <div className="border rounded p-4">
        <h3>Unique Visitors</h3>

        <p className="text-3xl font-bold">
          {uniqueVisitors}
        </p>
      </div>
    </div>
  );
}