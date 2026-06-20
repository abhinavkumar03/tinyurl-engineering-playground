import axios from "axios";

const api = axios.create({
  baseURL:
    import.meta.env.VITE_API_URL ||
    "http://localhost:8080/api/v1"
});

export const createShortUrl = async (url) => {
  const response = await api.post("/urls", {
    url
  });

  return response.data;
};

export default api;