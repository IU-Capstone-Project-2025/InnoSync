import Navbar from "../navbar/page_logged_in"
import styles from "./page.module.css"


const HomePage = () => {
  return (
    <div className={styles.container}>
      
      <Navbar />

      {/* Hero Section */}
      <main className={styles.hero}>
        <div className={styles.heroContent}>
          {/* Left Content */}
          <div className={styles.heroLeft}>
            <h1 className={styles.heroTitle}>
              One Platform
              <br />
              to <span className={styles.titleAccent}>Work</span> and
              <br />
              <span className={styles.titleAccent}>Sync</span> It All.
            </h1>

            <p className={styles.heroSubtitle}>
              Forget endless searching. The perfect match is one sync away.
            </p>

            <div className={styles.ctaButtons}>
              <button className={styles.primaryCta}>
                Get Started
              </button>
              <button className={styles.secondaryCta}>
                Learn more
              </button>
            </div>

            {/* University Logo */}
            <div className={styles.universityLogo}>
              <div className={styles.logoContainer}>
                <div className={styles.logoIcon}>
                  <span className={styles.logoIconText}>IU</span>
                </div>
                <div className={styles.logoInfo}>
                  <div className={styles.logoTitle}>INNOPOLIS</div>
                  <div className={styles.logoSubtitle}>UNIVERSITY</div>
                </div>
              </div>
            </div>
          </div>

          {/* Right Illustration */}
          <div className={styles.heroRight}>
            <div className={styles.illustrationContainer}>
              <img
                src="/landing_page.png"
                alt="Collaborative work illustration with people working together"
                className={styles.heroIllustration}
              />
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};




export default HomePage;

