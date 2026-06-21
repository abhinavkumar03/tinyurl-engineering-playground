import { Link, useLocation } from "react-router-dom";

export default function Navbar() {
  const location = useLocation();

  const isActive = (path) =>
    location.pathname === path;

  return (
    <nav className="navbar">
      <div className="navbar-brand">
        TinyURL
      </div>

      <div className="navbar-links">
        <Link
          to="/"
          className={
            isActive("/")
              ? "active"
              : ""
          }
        >
          Home
        </Link>

        <Link
          to="/dashboard"
          className={
            isActive("/dashboard")
              ? "active"
              : ""
          }
        >
          Dashboard
        </Link>
      </div>
    </nav>
  );
}