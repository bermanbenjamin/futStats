import AppModals from "@/components/modals";
import { SidebarProvider } from "@/components/ui/sidebar";
import { AppProviders } from "@/providers/app-provider";
import { ThemeProvider } from "@/providers/theme-provider";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "PitchStats",
  description:
    "PitchStats is a platform for soccer statistics with your friends.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt" suppressHydrationWarning={true}>
      <body className={inter.className}>
        <ThemeProvider
          attribute="class"
          defaultTheme="light"
          enableSystem
          disableTransitionOnChange
        >
          <SidebarProvider>
            <AppProviders>
              <AppModals />
              {children}
            </AppProviders>
          </SidebarProvider>
        </ThemeProvider>
      </body>
    </html>
  );
}
