import { useEffect, useState } from "react";

import {
  getDashboard,
} from "../services/analyticsService";

import TopURLsTable from "../components/TopURLsTable";

export default function DashboardPage() {
  const [dashboard, setDashboard] =
    useState(null);

  useEffect(() => {
    load();
  }, []);

  async function load() {
    const data = await getDashboard();

    setDashboard(data);
  }

  if (!dashboard) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-6">
      <h1>
        Analytics Dashboard
      </h1>

      <TopURLsTable
        urls={dashboard.top_urls}
      />
    </div>
  );
}