// src/app/authentication/layout.tsx

import React from "react";

export default function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      {/* Common layout wrapper for auth pages */}
      {children}
    </div>
  );
}