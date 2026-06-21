import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
import "./styles.css";
import "./components/Navbar.css";
import "./components/UrlForm.css";
import "./components/UrlResult.css";
import "./components/TopURLsTable.css";
import "./components/DailyClicksChart.css";
import "./components/AnalyticsCards.css";
import "./components/AnalyticsSummary.css";

ReactDOM.createRoot(
  document.getElementById("root")
).render(
  <React.StrictMode>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </React.StrictMode>
);