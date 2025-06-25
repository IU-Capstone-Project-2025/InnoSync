// components/Navbar.jsx
import styles from "./page.module.css"; // create this CSS module too
import React from "react";

const Navbar = () => {
  return (
    <nav className={styles.navbar}>
      <div className={styles.navContent}>
        {/* Logo */}
        <div className={styles.logo}>
          <span className={styles.logoText}>
            Inno<span className={styles.logoAccent}>Sync</span>
          </span>
        </div>

        {/* Navigation Links */}
        <div className={styles.navLinks}>
          <button className={styles.navLink}>
            <span>Find Project</span>
            <svg className={styles.chevron} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
              <path d="m6 9 6 6 6-6" />
            </svg>
          </button>

          <button className={styles.navLink}>
            <span>Find Talent</span>
            <svg className={styles.chevron} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
              <path d="m6 9 6 6 6-6" />
            </svg>
          </button>

          <button className={styles.navLink}>
            <span>About Us</span>
            <svg className={styles.chevron} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
              <path d="m6 9 6 6 6-6" />
            </svg>
          </button>
        </div>

        {/* Auth Buttons */}
        <div className={styles.authButtons}>
          <button className={styles.loginButton}>Login</button>
          <button className={styles.signupButton}>Sign Up</button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
