"use client";

import { Skeleton } from "@/components/ui/skeleton";
import { useGetPlayerById } from "@/http/player/use-player-service";
import { appRoutes } from "@/lib/routes";
import {
  Calendar,
  Footprints,
  Handshake,
  Shield,
  Square,
  Target,
} from "lucide-react";
import Link from "next/link";
import { use } from "react";
import StatsCard from "./components/StatsCard";
import StatsChart from "./components/StatsChart";

export default function PlayerPage({
  params,
}: {
  params: Promise<{ playerId: string }>;
}) {
  const { playerId } = use(params);
  const { data: player, isLoading } = useGetPlayerById(playerId);

  if (isLoading) {
    return (
      <div className="w-full space-y-6 p-6">
        <Skeleton className="h-8 w-48" />
        <Skeleton className="h-4 w-64" />
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {Array.from({ length: 7 }).map((_, i) => (
            <Skeleton key={i} className="h-16 w-full rounded-lg" />
          ))}
        </div>
        <Skeleton className="h-64 w-full" />
      </div>
    );
  }

  if (!player) return null;

  const allLeagues = [
    ...(player.owned_leagues ?? []),
    ...(player.member_of_leagues ?? []),
  ].filter(
    (league, index, self) => self.findIndex((l) => l.id === league.id) === index
  );

  return (
    <div className="w-full space-y-8 p-6">
      <div>
        <h1 className="text-2xl font-bold">{player.name}</h1>
        <p className="text-gray-500 text-sm">
          {player.position} · {player.email}
        </p>
      </div>

      <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        <StatsCard title="Goals" value={player.goals} icon={Target} />
        <StatsCard title="Assists" value={player.assists} icon={Handshake} />
        <StatsCard title="Matches" value={player.matches} icon={Calendar} />
        <StatsCard title="Disarms" value={player.disarms} icon={Shield} />
        <StatsCard title="Dribbles" value={player.dribbles} icon={Footprints} />
        <StatsCard
          title="Yellow Cards"
          value={player.yellow_cards}
          icon={Square}
        />
        <StatsCard title="Red Cards" value={player.red_cards} icon={Square} />
      </div>

      <StatsChart player={player} />

      {allLeagues.length > 0 ? (
        <div className="space-y-3">
          <h2 className="text-xl font-semibold">My Leagues</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-3">
            {allLeagues.map((league) => (
              <Link
                key={league.id}
                href={appRoutes.player.leagues.get(playerId, league.slug)}
                className="p-4 border rounded-lg hover:bg-gray-50 transition-colors"
              >
                <p className="font-medium">{league.name}</p>
                <p className="text-sm text-gray-500">
                  {league.members?.length ?? 0} members
                </p>
              </Link>
            ))}
          </div>
        </div>
      ) : (
        <div className="text-gray-500 text-center py-8 border rounded-lg">
          No leagues yet. Create one to get started!
        </div>
      )}
    </div>
  );
}
