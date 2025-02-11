"use client";

import { Icons } from "@/components/icons";
import { appRoutes } from "@/lib/routes";
import { useSessionStore } from "@/stores/session-store";
import { useRouter } from "next/navigation";
import StatsCard from "./components/StatsCard";
import StatsChart from "./components/StatsChart";

export default function PlayerPage() {
  const router = useRouter();
  const { player } = useSessionStore();

  if (!player) {
    router.push(appRoutes.auth.signIn);
    return;
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
