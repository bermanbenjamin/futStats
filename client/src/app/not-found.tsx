"use client";

import { appRoutes } from "@/lib/routes";
import { useSessionStore } from "@/stores/session-store";
import { redirect } from "next/navigation";

export default function Home() {
  const { player } = useSessionStore();

  if (player) {
    redirect(appRoutes.player.home(player.ID));
  } else {
    redirect(appRoutes.auth.signIn);
  }
}
