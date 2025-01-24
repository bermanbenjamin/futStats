import { getCookie } from "cookies-next";
import ky from "ky-universal";

export const api = ky.create({
  prefixUrl: process.env.NEXT_PUBLIC_SERVER_URL,
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 100000,
  retry: 2,
  hooks: {
    beforeRequest: [
      async (request) => {
        const token = await getCookie("token");
        if (token) {
          request.headers.set("Authorization", `Bearer ${token}`);
        }
      },
    ],
  },
});
