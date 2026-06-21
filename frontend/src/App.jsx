import {
  Routes,
  Route,
} from "react-router-dom";

import Navbar from "./components/Navbar";

import Home from "./pages/Home";
import DashboardPage from "./pages/DashboardPage";
import URLAnalyticsPage from "./pages/URLAnalyticsPage";

function App() {
  return (
    <>
      <Navbar />

      <Routes>
        <Route
          path="/"
          element={<Home />}
        />

        <Route
          path="/dashboard"
          element={<DashboardPage />}
        />

        <Route
          path="/analytics/:shortCode"
          element={<URLAnalyticsPage />}
        />
      </Routes>
    </>
  );
}

export default App;