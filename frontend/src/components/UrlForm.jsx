import { useState } from "react";

function UrlForm({ onSuccess }) {
  const [url, setUrl] =
    useState("");

  const [loading, setLoading] =
    useState(false);

  const [error, setError] =
    useState("");

  const validateURL = (value) => {
    try {
      new URL(value);
      return true;
    } catch {
      return false;
    }
  };

  const submit = async (e) => {
    e.preventDefault();

    setError("");

    const trimmedUrl = url.trim();

    if (!trimmedUrl) {
      setError(
        "Please enter a URL."
      );
      return;
    }

    if (
      !validateURL(trimmedUrl)
    ) {
      setError(
        "Please enter a valid URL including https://"
      );
      return;
    }

    try {
      setLoading(true);

      const {
        createShortUrl,
      } = await import(
        "../services/api"
      );

      const response =
        await createShortUrl(
          trimmedUrl
        );

      onSuccess(response);

      setUrl("");
    } catch (error) {
      setError(
        error?.response?.data
          ?.error ||
          "Unable to create short URL."
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <form
      onSubmit={submit}
      className="url-form"
    >
      <label
        htmlFor="url-input"
        className="form-label"
      >
        Enter destination URL
      </label>

      <div className="url-form-row">

        <input
          id="url-input"
          type="text"
          value={url}
          placeholder="https://example.com"
          onChange={(e) =>
            setUrl(
              e.target.value
            )
          }
          className={`url-input ${
            error
              ? "input-error"
              : ""
          }`}
        />

        <button
          type="submit"
          disabled={loading}
          className="submit-btn"
        >
          {loading
            ? "Creating..."
            : "Shorten URL"}
        </button>

      </div>

      {error && (
        <p className="error-text">
          {error}
        </p>
      )}
    </form>
  );
}

export default UrlForm;