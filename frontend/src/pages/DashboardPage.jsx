import { useEffect, useState } from "react";

import { getDashboard }
  from "../services/analyticsService";

import TopURLsTable
  from "../components/TopURLsTable";

export default function DashboardPage() {

  const [dashboard, setDashboard] =
    useState(null);

  const [loading, setLoading] =
    useState(true);

  const [error, setError] =
    useState("");

  useEffect(() => {
    load();
  }, []);

  async function load() {
    try {
      setLoading(true);

      const data =
        await getDashboard();

      setDashboard(data);

    } catch (err) {

      setError(
        "Failed to load dashboard."
      );

    } finally {
      setLoading(false);
    }
  }

  if (loading) {
    return (
      <div className="page-container">
        <DashboardSkeleton />
      </div>
    );
  }

  if (error) {
    return (
      <div className="page-container">
        <div className="error-state">
          {error}
        </div>
      </div>
    );
  }

  const urls =
    dashboard?.top_urls || [];

  const totalUrls =
    urls.length;

  const totalClicks =
    urls.reduce(
      (sum, item) =>
        sum + item.clicks,
      0
    );

  const bestUrl =
    urls.length > 0
      ? urls[0]
      : null;

  return (
    <div className="page-container">

      <div className="dashboard-header">

        <h1>
          Analytics Dashboard
        </h1>

        <p>
          Monitor URL performance
          and explore click data.
        </p>

      </div>

      <div className="stats-grid">

        <div className="stat-card">

          <span>
            Tracked URLs
          </span>

          <h2>
            {totalUrls}
          </h2>

        </div>

        <div className="stat-card">

          <span>
            Total Clicks
          </span>

          <h2>
            {totalClicks}
          </h2>

        </div>

        <div className="stat-card">

          <span>
            Top Performer
          </span>

          <h2>
            {bestUrl
              ? bestUrl.short_code
              : "--"}
          </h2>

        </div>

      </div>

      <section className="dashboard-section">

        <div className="section-header-row">

          <div>

            <h2>
              Top Performing URLs
            </h2>

            <p>
              URLs ranked by click
              activity.
            </p>

          </div>

        </div>

        {urls.length === 0 ? (
          <div className="empty-state">

            <h3>
              No URL activity yet
            </h3>

            <p>
              Create your first
              short URL to begin
              collecting analytics.
            </p>

          </div>
        ) : (
          <TopURLsTable
            urls={urls}
          />
        )}

      </section>

    </div>
  );
}

function DashboardSkeleton() {
  return (
    <>
      <div className="skeleton-header" />

      <div className="stats-grid">

        <div className="skeleton-card" />
        <div className="skeleton-card" />
        <div className="skeleton-card" />

      </div>

      <div className="skeleton-table" />
    </>
  );
}