import Link from "next/link";
import styles from "./page.module.css";

const NavbarUser = () => {
  return (
    <nav className={styles.navbar}>
      <div className={styles.navContent}>
        {/* Logo */}
        <div className={styles.logo}>
          <span className={styles.logoText}>
            <Link href="/components/home" className={styles.navLink}>  
            Inno<span className={styles.logoAccent}>Sync</span> 
            </Link>
          </span>
        </div>

        {/* Navigation Links */}
        <div className={styles.navLinks}>
          <Link href="/components/projects" className={styles.navLink}>
            Find Project
          </Link>
          <Link href="/components/talent" className={styles.navLink}>
            Find Talent
          </Link>
          <Link href="/components/about" className={styles.navLink}>
            About Us
          </Link>
        </div>

        {/* Logged-in User UI */}
        <div className={styles.userInfo}>
          <span className={styles.userItem}>
            <svg className={styles.bellIcon} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
              <path d="M18 8a6 6 0 10-12 0c0 7-3 9-3 9h18s-3-2-3-9" />
              <path d="M13.73 21a2 2 0 01-3.46 0" />
            </svg>
            Recruiter
            <svg className={styles.arrowIcon} viewBox="0 0 24 24" fill="none" stroke="green" strokeWidth="3">
              <circle cx="12" cy="12" r="10" />
              <path d="M10 8l6 6-6 6" strokeLinecap="round" strokeLinejoin="round" transform="scale(0.8) translate(3, 3)" />
            </svg>
          </span>

          <span className={styles.userItem}>
            A.Alimi
            <svg className={styles.arrowIcon} viewBox="0 0 24 24" fill="none" stroke="green" strokeWidth="3">
              <circle cx="12" cy="12" r="10" />
              <path d="M6 9l6 6 6-6" strokeLinecap="round" strokeLinejoin="round" transform="scale(0.8) translate(3, 3)" />
            </svg>
          </span>
        </div>
      </div>
    </nav>
  );
};

export default NavbarUser;
