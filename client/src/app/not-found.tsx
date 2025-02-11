"use client";

import { appRoutes } from "@/lib/routes";
import { useSessionStore } from "@/stores/session-store";
import { getCookie } from "cookies-next";
import { redirect } from "next/navigation";

export default function NotFoundPage() {
  const token = getCookie("token");
  const { player } = useSessionStore();

  if (token) {
    console.log(player);
    redirect(appRoutes.player.home(player!.ID));
  } else {
    redirect(appRoutes.auth.signIn);
  }
}
