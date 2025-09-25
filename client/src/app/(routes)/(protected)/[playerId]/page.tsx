"use client";

import { Icons } from "@/components/icons";
import { appRoutes } from "@/lib/routes";
import { useSessionStore } from "@/stores/session-store";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import StatsCard from "./components/StatsCard";
import StatsChart from "./components/StatsChart";

// Force dynamic rendering
export const dynamic = "force-dynamic";

export default function PlayerPage() {
  const router = useRouter();
  const { player } = useSessionStore();

  useEffect(() => {
    if (!player) {
      router.push(appRoutes.auth.signIn);
    }
  }, [player, router]);

  if (!player) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-center">
          <h1 className="text-2xl font-bold mb-4">Loading...</h1>
          <p className="text-gray-600">Redirecting to sign in...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="w-full">
      <h1 className=" text-xl mb-2">
        Bem vindo de volta,
        <span className="font-semibold text-foreground"> {player.name}.</span>
      </h1>
      <h2 className=" text-sm font-light text-foreground">
        {" "}
        Aqui estão suas estatisticas da temporada:{" "}
      </h2>

      <div className="grid grid-cols-2 gap-4 mt-4">
        <StatsCard title="Gols" value={player.goals} icon={Icons.trophy} />
        <StatsCard
          title="Assistências"
          value={player.assists}
          icon={Icons.trophy}
        />
        <StatsCard
          title="Dribles"
          value={player.dribbles}
          icon={Icons.trophy}
        />
        <StatsCard
          title="Desarmes"
          value={player.disarms}
          icon={Icons.trophy}
        />
        <StatsCard title="Jogos" value={player.matches} icon={Icons.trophy} />
      </div>

      <div className="mt-4 w-full h-full">
        <StatsChart player={player} />
      </div>
    </div>
  );
}
