import { Link } from "react-router-dom";

export default function TopURLsTable({
  urls,
}) {
  if (!urls?.length) {
    return null;
  }

  return (
    <>
      {/* Desktop Table */}

      <div className="table-wrapper">

        <table className="analytics-table">

          <thead>

            <tr>
              <th>Short Code</th>
              <th>Original URL</th>
              <th>Clicks</th>
              <th>Action</th>
            </tr>

          </thead>

          <tbody>

            {urls.map((url) => (
              <tr
                key={url.short_code}
              >
                <td>
                  <span className="code-pill">
                    {url.short_code}
                  </span>
                </td>

                <td>
                  <a
                    href={url.original_url}
                    target="_blank"
                    rel="noreferrer"
                    className="url-cell"
                  >
                    {url.original_url}
                  </a>
                </td>

                <td>
                  <span className="click-badge">
                    {url.clicks}
                  </span>
                </td>

                <td>

                  <Link
                    to={`/analytics/${url.short_code}`}
                    className="table-action-btn"
                  >
                    View Analytics
                  </Link>

                </td>
              </tr>
            ))}

          </tbody>

        </table>

      </div>

      {/* Mobile Cards */}

      <div className="mobile-url-list">

        {urls.map((url) => (
          <div
            key={url.short_code}
            className="mobile-url-card"
          >

            <div className="mobile-url-header">

              <span className="code-pill">
                {url.short_code}
              </span>

              <span className="click-badge">
                {url.clicks} Clicks
              </span>

            </div>

            <a
              href={url.original_url}
              target="_blank"
              rel="noreferrer"
              className="mobile-url-link"
            >
              {url.original_url}
            </a>

            <Link
              to={`/analytics/${url.short_code}`}
              className="table-action-btn"
            >
              View Analytics
            </Link>

          </div>
        ))}

      </div>
    </>
  );
}