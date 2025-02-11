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
        <div className="bg-neutral-100 dark:bg-neutral-800 p-4 rounded-lg">
          <h3 className="text-lg font-medium text-neutral-900 dark:text-neutral-100">
            Gols
          </h3>
          <p className="text-2xl font-semibold text-indigo-900 dark:text-indigo-100">
            {player.goals}
          </p>
        </div>
        <div className="bg-neutral-100 dark:bg-neutral-800 p-4 rounded-lg">
          <h3 className="text-lg font-medium text-neutral-900 dark:text-neutral-100">
            Assistências
          </h3>
          <p className="text-2xl font-semibold text-indigo-900 dark:text-indigo-100">
            {player.assists}
          </p>
        </div>
        <div className="bg-neutral-100 dark:bg-neutral-800 p-4 rounded-lg">
          <h3 className="text-lg font-medium text-neutral-900 dark:text-neutral-100">
            Dribles
          </h3>
          <p className="text-2xl font-semibold text-indigo-900 dark:text-indigo-100">
            {player.dribbles}
          </p>
        </div>
        <div className="bg-neutral-100 dark:bg-neutral-800 p-4 rounded-lg">
          <h3 className="text-lg font-medium text-neutral-900 dark:text-neutral-100">
            Desarmes
          </h3>
          <p className="text-2xl font-semibold text-indigo-900 dark:text-indigo-100">
            {player.disarms}
          </p>
        </div>
        <div className="bg-neutral-100 dark:bg-neutral-800 p-4 rounded-lg">
          <h3 className="text-lg font-medium text-neutral-900 dark:text-neutral-100">
            Jogos
          </h3>
          <p className="text-2xl font-semibold text-indigo-900 dark:text-indigo-100">
            {player.matches}
          </p>
        </div>
      </div>
    </div>
  );
}
