"use client";

import { AddPlayerModal } from "@/components/modals/add-player";
import { RecordMatchModal } from "@/components/modals/record-match";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { useGetMatchesByLeague } from "@/http/match/use-match-service";
import { useGetLeagueService } from "@/http/league/use-league-service";
import { useQueryState } from "nuqs";
import { use } from "react";
import LeagueOverview from "./components/league-overview";
import PlayersSection from "./components/players-section";
import TransferHistorySection from "./components/transfer-history";

export default function LeaguePage({
  params,
}: {
  params: Promise<{ playerId: string; leagueId: string }>;
}) {
  const { leagueId } = use(params);
  const { data: league, isLoading } = useGetLeagueService(leagueId);
  const { data: matches } = useGetMatchesByLeague(leagueId);
  const [, setRecordMatch] = useQueryState("record-match");

  if (isLoading) {
    return (
      <div className="w-full space-y-6 p-6">
        <Skeleton className="h-8 w-64" />
        <div className="grid grid-cols-2 gap-4">
          <Skeleton className="h-32 w-full" />
          <Skeleton className="h-32 w-full" />
        </div>
        <Skeleton className="h-64 w-full" />
      </div>
    );
  }

  if (!league) {
    return (
      <div className="w-full p-6 text-gray-500 text-center">
        League not found.
      </div>
    );
  }

  return (
    <div className="w-full space-y-8 p-6">
      <div className="flex items-start justify-between">
        <LeagueOverview
          name={league.name}
          ownerName={league.owner?.name ?? ""}
          createdAt={league.createdAt}
          membersLength={league.members?.length ?? 0}
        />
        <Button onClick={() => setRecordMatch("true")} className="shrink-0">
          Record Match
        </Button>
      </div>

      {matches && matches.length > 0 && (
        <div className="space-y-2">
          <h2 className="text-xl font-semibold">Match History</h2>
          <div className="space-y-2">
            {matches.map((match) => (
              <div key={match.id} className="p-3 border rounded-lg flex justify-between items-center">
                <span className="text-sm">
                  {new Date(match.date).toLocaleDateString("pt-BR", {
                    day: "2-digit",
                    month: "2-digit",
                    year: "numeric",
                    hour: "2-digit",
                    minute: "2-digit",
                  })}
                </span>
                <span className="text-sm text-gray-500">
                  {match.events?.length ?? 0} events
                </span>
              </div>
            ))}
          </div>
        </div>
      )}

      <PlayersSection players={league.members ?? []} />
      <TransferHistorySection />
      <AddPlayerModal />
      <RecordMatchModal
        leagueId={league.id}
        leagueSlug={leagueId}
        players={league.members ?? []}
      />
    </div>
  );
}
