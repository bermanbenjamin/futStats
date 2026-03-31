"use client";

import { AppSidebar } from "@/components/sidebar";
import { getCookie } from "cookies-next";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function ProtectedLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();

  useEffect(() => {
    const token = getCookie("token");
    if (!token) {
      router.replace("/auth/sign-in");
    }
  }, [router]);

  return (
    <main className="min-h-screen w-full bg-white">
      <section className="flex size-full p-8">
        <AppSidebar />
        {children}
      </section>
    </main>
  );
}
