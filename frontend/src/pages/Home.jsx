import { useState } from "react";
import { Link } from "react-router-dom";

import UrlForm from "../components/UrlForm";
import UrlResult from "../components/UrlResult";

function Home() {
  const [result, setResult] = useState(null);

  return (
    <main className="home-page">

      {/* Hero */}

      <section className="hero-section">

        <span className="hero-badge">
          Production-Inspired URL Platform
        </span>

        <h1 className="hero-title">
          Shorten URLs.
          <br />
          Track Performance.
          <br />
          Understand Engagement.
        </h1>

        <p className="hero-description">
          TinyURL Engineering Playground
          demonstrates URL shortening,
          click tracking, analytics,
          and dashboard reporting in a
          clean SaaS experience.
        </p>

      </section>

      {/* URL Creator */}

      <section className="url-creator-section">

        <div className="card">

          <div className="section-header">
            <h2>Create Short URL</h2>

            <p>
              Paste a URL and generate
              a shareable short link.
            </p>
          </div>

          <UrlForm
            onSuccess={(data) =>
              setResult(data)
            }
          />

        </div>

      </section>

      {/* Result */}

      {result && (
        <section className="result-section">

          <UrlResult
            result={result}
          />

        </section>
      )}

      {/* Features */}

      <section className="features-section">

        <div className="feature-card">
          <h3>
            Fast URL Shortening
          </h3>

          <p>
            Generate clean and
            shareable links instantly.
          </p>
        </div>

        <div className="feature-card">
          <h3>
            Click Analytics
          </h3>

          <p>
            Measure engagement using
            total clicks and visitor
            metrics.
          </p>
        </div>

        <div className="feature-card">
          <h3>
            Dashboard Insights
          </h3>

          <p>
            Track top-performing URLs
            from a centralized
            dashboard.
          </p>
        </div>

      </section>

      {/* CTA */}

      <section className="analytics-cta">

        <div className="card">

          <h2>
            Explore Analytics
          </h2>

          <p>
            View top URLs and
            performance metrics in
            the analytics dashboard.
          </p>

          <Link
            to="/dashboard"
            className="primary-btn"
          >
            Open Dashboard
          </Link>

        </div>

      </section>

    </main>
  );
}

export default Home;