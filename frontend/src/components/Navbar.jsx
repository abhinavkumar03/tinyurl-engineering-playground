import { Link, useLocation } from "react-router-dom";
import { useState } from "react";

export default function Navbar() {
  const location = useLocation();

  const [mobileOpen, setMobileOpen] =
    useState(false);

  const navItems = [
    {
      label: "Home",
      path: "/",
    },
    {
      label: "Dashboard",
      path: "/dashboard",
    },
  ];

  const isActive = (path) =>
    location.pathname === path;

  return (
    <header className="navbar">
      <div className="navbar-container">

        <Link
          to="/"
          className="navbar-logo"
        >
          <div className="logo-title">
            TinyURL
          </div>

          <div className="logo-subtitle">
            Engineering Playground
          </div>
        </Link>

        <nav className="desktop-nav">
          {navItems.map((item) => (
            <Link
              key={item.path}
              to={item.path}
              className={`nav-link ${
                isActive(item.path)
                  ? "nav-link-active"
                  : ""
              }`}
            >
              {item.label}
            </Link>
          ))}
        </nav>

        <button
          className="mobile-menu-btn"
          onClick={() =>
            setMobileOpen(!mobileOpen)
          }
        >
          ☰
        </button>

      </div>

      {mobileOpen && (
        <div className="mobile-nav">
          {navItems.map((item) => (
            <Link
              key={item.path}
              to={item.path}
              onClick={() =>
                setMobileOpen(false)
              }
              className={`mobile-nav-link ${
                isActive(item.path)
                  ? "nav-link-active"
                  : ""
              }`}
            >
              {item.label}
            </Link>
          ))}
        </div>
      )}
    </header>
  );
}