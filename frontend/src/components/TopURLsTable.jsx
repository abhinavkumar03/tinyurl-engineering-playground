export default function TopURLsTable({
  urls,
}) {
  return (
    <table className="w-full border">
      <thead>
        <tr>
          <th>Short Code</th>
          <th>Original URL</th>
          <th>Clicks</th>
        </tr>
      </thead>

      <tbody>
        {urls.map((url) => (
          <tr key={url.short_code}>
            <td>{url.short_code}</td>

            <td>{url.original_url}</td>

            <td>{url.clicks}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}