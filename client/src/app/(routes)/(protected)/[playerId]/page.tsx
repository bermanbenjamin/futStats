"use client";

import { appRoutes } from "@/lib/routes";
import { useSessionStore } from "@/stores/session-store";
import { useRouter } from "next/navigation";

export default function PlayerPage() {
  const router = useRouter();
  const { player } = useSessionStore();

  if (!player) {
    router.push(appRoutes.auth.signIn);
    return;
  }

  return (
    <div>
      <h1>Player Page</h1>
      <p>Player ID: {player.name}</p>
    </div>
  );
}
