import { useState } from "react";
import { Link } from "react-router-dom";

function UrlResult({ result }) {
  const [copied, setCopied] =
    useState(false);

  if (!result) {
    return null;
  }

  const copy = async () => {
    try {
      await navigator.clipboard.writeText(
        result.short_url
      );

      setCopied(true);

      setTimeout(() => {
        setCopied(false);
      }, 2000);

    } catch {
      console.error(
        "Failed to copy URL"
      );
    }
  };

  return (
    <div className="success-card">

      <div className="success-header">

        <div className="success-icon">
          ✓
        </div>

        <div>
          <h3>
            URL Created Successfully
          </h3>

          <p>
            Your shortened link is
            ready to share.
          </p>
        </div>

      </div>

      <div className="short-url-section">

        <label>
          Short URL
        </label>

        <div className="short-url-box">

          <a
            href={result.short_url}
            target="_blank"
            rel="noreferrer"
          >
            {result.short_url}
          </a>

          <button
            onClick={copy}
            className="copy-btn"
          >
            {copied
              ? "Copied!"
              : "Copy"}
          </button>

        </div>

      </div>

      <div className="original-url-section">

        <label>
          Original URL
        </label>

        <div className="original-url-box">
          {result.original_url}
        </div>

      </div>

      <div className="result-actions">

        <Link
          to={`/analytics/${result.short_code}`}
          className="primary-btn"
        >
          View Analytics
        </Link>

      </div>

    </div>
  );
}

export default UrlResult;