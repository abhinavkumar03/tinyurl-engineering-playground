import { useState } from "react";

function UrlForm({ onSuccess }) {
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);

  const submit = async (e) => {
    e.preventDefault();

    if (!url.trim()) {
      return;
    }

    try {
      setLoading(true);

      const { createShortUrl } = await import(
        "../services/api"
      );

      const response = await createShortUrl(url);

      onSuccess(response);

      setUrl("");
    } catch (error) {
      alert(
        error?.response?.data?.error ||
          "Failed to create URL"
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={submit} className="url-form">
      <input
        type="url"
        placeholder="https://example.com"
        value={url}
        onChange={(e) =>
          setUrl(e.target.value)
        }
        required
      />

      <button
        type="submit"
        disabled={loading}
      >
        {loading
          ? "Creating..."
          : "Shorten URL"}
      </button>
    </form>
  );
}

export default UrlForm;