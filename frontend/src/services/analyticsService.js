import axios from "axios";

const API = import.meta.env.VITE_API_URL;

export async function getDashboard() {
  const { data } = await axios.get(
    `${API}/analytics/dashboard`
  );

  return data;
}

export async function getTopURLs(limit = 10) {
  const { data } = await axios.get(
    `${API}/analytics/top-urls?limit=${limit}`
  );

  return data;
}

export async function getURLAnalytics(shortCode) {
  const { data } = await axios.get(
    `${API}/analytics/${shortCode}`
  );

  return data;
}

export async function getDailyClicks(shortCode) {
  const { data } = await axios.get(
    `${API}/analytics/${shortCode}/daily`
  );

  return data;
}