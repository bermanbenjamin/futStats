"use client";

import { Toaster } from "@/components/ui/sonner";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { NuqsAdapter } from "nuqs/adapters/next/app";

interface Props {
  children: React.ReactNode;
}

const AppProviders = ({ children }: Props) => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
        refetchOnWindowFocus: false,
      },
    },
  });

  return (
    <QueryClientProvider client={queryClient}>
      <NuqsAdapter>
        <Toaster />
        {children}
      </NuqsAdapter>
    </QueryClientProvider>
  );
};

export { AppProviders };
