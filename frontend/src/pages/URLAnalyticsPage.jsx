import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import {
  getURLAnalytics,
  getDailyClicks,
} from "../services/analyticsService";

import AnalyticsSummary from "../components/AnalyticsSummary";
import DailyClicksChart from "../components/DailyClicksChart";

export default function URLAnalyticsPage() {

  const { shortCode } = useParams();

  const [analytics, setAnalytics] =
    useState(null);

  const [dailyClicks, setDailyClicks] =
    useState([]);

  useEffect(() => {
    load();
  }, [shortCode]);

  async function load() {

    const analyticsData =
      await getURLAnalytics(shortCode);

    const clicksData =
      await getDailyClicks(shortCode);

    setAnalytics(
      analyticsData,
    );

    setDailyClicks(
      clicksData,
    );
  }

  if (!analytics) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-6">

      <AnalyticsSummary
        analytics={analytics}
      />

      <DailyClicksChart
        data={dailyClicks}
      />

    </div>
  );
}