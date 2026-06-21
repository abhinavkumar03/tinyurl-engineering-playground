import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import {
  getURLAnalytics,
  getDailyClicks,
} from "../services/analyticsService";

import AnalyticsCards from "../components/AnalyticsCards";
import AnalyticsSummary from "../components/AnalyticsSummary";
import DailyClicksChart from "../components/DailyClicksChart";

export default function URLAnalyticsPage() {
  const { shortCode } = useParams();

  const [analytics, setAnalytics] = useState(null);
  const [dailyClicks, setDailyClicks] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    loadAnalytics();
  }, [shortCode]);

  async function loadAnalytics() {
    try {
      setLoading(true);
      setError("");

      const [analyticsData, clicksData] =
        await Promise.all([
          getURLAnalytics(shortCode),
          getDailyClicks(shortCode),
        ]);

      setAnalytics(analyticsData);
      setDailyClicks(clicksData);
    } catch (err) {
      setError(
        "Unable to load analytics. Please try again."
      );
    } finally {
      setLoading(false);
    }
  }

  if (loading) {
    return (
      <div className="page-container">
        <AnalyticsSkeleton />
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

  return (
    <div className="page-container">
      {/* Header */}

      <div className="analytics-header">
        <h1>URL Analytics</h1>

        <p>
          Performance insights for short code{" "}
          <strong>
            {analytics.short_code}
          </strong>
        </p>
      </div>

      {/* KPI Cards */}

      <AnalyticsCards
        totalClicks={
          analytics.total_clicks
        }
        uniqueVisitors={
          analytics.unique_visitors
        }
      />

      {/* Chart + Details */}

      <div className="analytics-layout">
        {/* Chart */}

        <div className="analytics-panel">
          <h2>
            Daily Click Trend
          </h2>

          {dailyClicks.length === 0 ? (
            <div className="empty-state">
              <h3>
                No click activity yet
              </h3>

              <p>
                Click data will appear
                once this URL receives
                traffic.
              </p>
            </div>
          ) : (
            <DailyClicksChart
              data={dailyClicks}
            />
          )}
        </div>

        {/* URL Details */}

        <div className="analytics-panel">
          <AnalyticsSummary
            analytics={analytics}
          />
        </div>
      </div>

      {/* Insights */}

      <div className="analytics-panel insights-panel">
        <h2>
          Analytics Insights
        </h2>

        <div className="insights-list">
          <div className="insight-item">
            <strong>
              Total Clicks
            </strong>

            <span>
              {
                analytics.total_clicks
              }
            </span>
          </div>

          <div className="insight-item">
            <strong>
              Unique Visitors
            </strong>

            <span>
              {
                analytics.unique_visitors
              }
            </span>
          </div>

          <div className="insight-item">
            <strong>
              Engagement Ratio
            </strong>

            <span>
              {analytics.unique_visitors
                ? (
                    analytics.total_clicks /
                    analytics.unique_visitors
                  ).toFixed(1)
                : "0"}
            </span>
          </div>

          <div className="insight-item">
            <strong>
              Traffic Status
            </strong>

            <span>
              {analytics.total_clicks >
              0
                ? "Receiving Traffic"
                : "No Traffic Yet"}
            </span>
          </div>
        </div>
      </div>
    </div>
  );
}

function AnalyticsSkeleton() {
  return (
    <>
      <div className="skeleton-header" />

      <div className="stats-grid">
        <div className="skeleton-card" />
        <div className="skeleton-card" />
      </div>

      <div className="analytics-layout">
        <div className="skeleton-table" />

        <div className="skeleton-table" />
      </div>
    </>
  );
}