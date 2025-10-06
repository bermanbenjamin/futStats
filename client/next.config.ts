import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "standalone",
  // Remove basePath - it's not needed for your use case
  // Instead, handle routing through your app structure
  async redirects() {
    return [
      {
        source: "/",
        destination: "/auth/sign-in",
        permanent: false,
      },
    ];
  },
};

export default nextConfig;
