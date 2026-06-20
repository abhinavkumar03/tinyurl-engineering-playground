function UrlResult({ result }) {
  if (!result) {
    return null;
  }

  const copy = async () => {
    await navigator.clipboard.writeText(
      result.short_url
    );

    alert("Copied");
  };

  return (
    <div className="result-card">
      <h3>Short URL Created</h3>

      <p>
        <strong>Original:</strong>
      </p>

      <a
        href={result.original_url}
        target="_blank"
        rel="noreferrer"
      >
        {result.original_url}
      </a>

      <p>
        <strong>Short URL:</strong>
      </p>

      <a
        href={result.short_url}
        target="_blank"
        rel="noreferrer"
      >
        {result.short_url}
      </a>

      <button onClick={copy}>
        Copy URL
      </button>
    </div>
  );
}

export default UrlResult;