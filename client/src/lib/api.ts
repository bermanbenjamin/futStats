import { getCookie } from "cookies-next";
import ky from "ky-universal";
import { redirect } from "next/navigation";

// Constants for configuration
const API_CONFIG = {
  baseUrl: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080",
  timeout: 100000,
  retries: 2,
  authTokenCookie: "token",
} as const;

interface ApiError extends Error {
  response?: {
    status: number;
  };
}

// Pure error handler with explicit side effects
const handleApiError = async (error: ApiError): Promise<never> => {
  // Explicit side effect: logging
  console.error("API Error:", {
    status: error.response?.status,
    message: error.message,
  });

  // Explicit side effect: authentication redirect
  if (error.response?.status === 401) {
    redirect("/auth/sign-in");
  }

  // Re-throw to maintain error chain
  throw error;
};

export const api = ky.create({
  prefixUrl: API_CONFIG.baseUrl + "/api",
  headers: {
    "Content-Type": "application/json",
  },
  timeout: API_CONFIG.timeout,
  retry: API_CONFIG.retries,
  hooks: {
    beforeRequest: [
      async (request) => {
        const token = await getCookie(API_CONFIG.authTokenCookie);
        if (token && !request.url.includes("auth")) {
          request.headers.set("Authorization", `Bearer ${token}`);
        }
      },
    ],
    afterResponse: [
      async (request, options, response) => {
        return response;
      },
    ],
    beforeError: [
      async (error) => {
        return handleApiError(error);
      },
    ],
  },
});
