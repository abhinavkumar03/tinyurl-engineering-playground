import { useState } from "react";
import UrlForm from "../components/UrlForm";
import UrlResult from "../components/UrlResult";

function Home() {
  const [result, setResult] = useState(null);

  return (
    <div className="container">
      <h1>
        TinyURL Engineering Playground
      </h1>

      <p>
        Production-inspired URL shortening
        service.
      </p>

      <UrlForm
        onSuccess={(data) =>
          setResult(data)
        }
      />

      <UrlResult result={result} />
    </div>
  );
}

export default Home;