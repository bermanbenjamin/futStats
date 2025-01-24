import { NextRequest, NextResponse } from "next/server";
import { appRoutes } from "./lib/routes";

export default function authMiddleware(request: NextRequest) {
  const requestHeaders = new Headers(request.headers);
  const token = requestHeaders.get("authorization");
  if (!token && !request.url.includes("auth")) {
    return NextResponse.redirect(appRoutes.auth.login);
  }
}
