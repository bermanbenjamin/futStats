"use client";

import { AppSidebar } from "@/components/sidebar";

export default function ProtectedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
  sheet: React.ReactNode;
}>) {
  return (
    <>
      <main className="min-h-screen w-full bg-white">
        <section className="flex size-full">
          <AppSidebar />
          {children}
        </section>
      </main>
    </>
  );
}
