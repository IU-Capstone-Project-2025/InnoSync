import { ReactNode } from 'react';
import Image from "next/image";
import styles from "./page.module.css";


export default function AuthLayout({ children }: { children: ReactNode }) {
    return (
    <div className={styles.parentContainer}>
      <div className={styles.backgroundContainer}>
        <Image
          src="/auth.svg"
          width={899}
          height={1080}
          className={styles.auth_logo}
          alt="auth decoration"
        />
      </div>
      <main className={styles.authContent}>
          {children}
        </main>
    </div>
  );
}