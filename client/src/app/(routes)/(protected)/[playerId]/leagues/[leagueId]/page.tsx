"use client";

import { AddPlayerModal } from "@/components/modals/add-player";
import { CreateSeasonModal } from "@/components/modals/create-season";
import { RecordMatchModal } from "@/components/modals/record-match";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import { useGetMatchesByLeague } from "@/http/match/use-match-service";
import { useGetLeagueService } from "@/http/league/use-league-service";
import { useGetSeasonsByLeague } from "@/http/season/use-season-service";
import { useSessionStore } from "@/stores/session-store";
import { useQueryState } from "nuqs";
import { use, useState } from "react";
import Link from "next/link";
import LeagueOverview from "./components/league-overview";
import PlayersSection from "./components/players-section";

export default function LeaguePage({
  params,
}: {
  params: Promise<{ playerId: string; leagueId: string }>;
}) {
  const { playerId, leagueId } = use(params);
  const [, setRecordMatch] = useQueryState("record-match");
  const [createSeasonOpen, setCreateSeasonOpen] = useState(false);

  const { player: sessionPlayer } = useSessionStore();
  const { data: league, isLoading } = useGetLeagueService(leagueId);
  const { data: matches } = useGetMatchesByLeague(leagueId);
  const { data: seasons } = useGetSeasonsByLeague(leagueId);

  const isOwner = league?.ownerId === sessionPlayer?.id;
  const activeSeason = seasons?.find((s) => s.status === "active");
  const pastSeasons = seasons?.filter((s) => s.status === "finished") ?? [];

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
        Liga não encontrada.
      </div>
    );
  }

  return (
    <div className="w-full space-y-8 p-6">
      {/* Header */}
      <div className="flex items-start justify-between">
        <LeagueOverview
          name={league.name}
          ownerName={league.owner?.name ?? ""}
          createdAt={league.createdAt}
          membersLength={league.members?.length ?? 0}
        />
        <Button onClick={() => setRecordMatch("true")} className="shrink-0">
          Registrar Partida
        </Button>
      </div>

      {/* Active season banner */}
      {activeSeason ? (
        <div className="flex items-center justify-between p-4 rounded-lg border border-green-200 bg-green-50 dark:bg-green-950/20 dark:border-green-800">
          <div className="flex items-center gap-3">
            <Badge className="bg-green-500 hover:bg-green-600">Ativa</Badge>
            <div>
              <p className="font-medium text-sm">{activeSeason.year}</p>
              <p className="text-xs text-muted-foreground">
                {activeSeason.init_date
                  ? new Date(activeSeason.init_date).toLocaleDateString("pt-BR")
                  : "—"}{" "}
                →{" "}
                {activeSeason.end_date
                  ? new Date(activeSeason.end_date).toLocaleDateString("pt-BR")
                  : "—"}
              </p>
            </div>
          </div>
          <Link
            href={`/${playerId}/leagues/${leagueId}/seasons/${activeSeason.id}`}
          >
            <Button variant="outline" size="sm">
              Ver Standings
            </Button>
          </Link>
        </div>
      ) : (
        <div className="flex items-center justify-between p-4 rounded-lg border border-dashed">
          <p className="text-sm text-muted-foreground">
            Nenhuma temporada ativa
          </p>
          {isOwner && (
            <Button
              variant="outline"
              size="sm"
              onClick={() => setCreateSeasonOpen(true)}
            >
              Iniciar Temporada
            </Button>
          )}
        </div>
      )}

      {/* Match history */}
      {matches && matches.length > 0 && (
        <div className="space-y-2">
          <h2 className="text-xl font-semibold">Histórico de Partidas</h2>
          <div className="space-y-2">
            {matches.map((match) => (
              <div
                key={match.id}
                className="p-3 border rounded-lg flex justify-between items-center"
              >
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
                  {match.events?.length ?? 0} eventos
                </span>
              </div>
            ))}
          </div>
        </div>
      )}

      <PlayersSection players={league.members ?? []} />

      {/* Past seasons */}
      {pastSeasons.length > 0 && (
        <div className="space-y-3">
          <h2 className="text-xl font-semibold">Temporadas Anteriores</h2>
          <div className="space-y-2">
            {pastSeasons.map((season) => (
              <Link
                key={season.id}
                href={`/${playerId}/leagues/${leagueId}/seasons/${season.id}`}
                className="flex items-center justify-between p-3 border rounded-lg hover:bg-accent transition-colors"
              >
                <div>
                  <p className="font-medium text-sm">{season.year}</p>
                  <p className="text-xs text-muted-foreground">
                    {season.init_date
                      ? new Date(season.init_date).toLocaleDateString("pt-BR")
                      : "—"}{" "}
                    →{" "}
                    {season.end_date
                      ? new Date(season.end_date).toLocaleDateString("pt-BR")
                      : "—"}
                  </p>
                </div>
                <span className="text-sm text-muted-foreground">Ver →</span>
              </Link>
            ))}
          </div>
        </div>
      )}

      <AddPlayerModal />
      <CreateSeasonModal
        leagueSlug={leagueId}
        open={createSeasonOpen}
        onOpenChange={setCreateSeasonOpen}
      />
      <RecordMatchModal
        leagueId={league.id}
        leagueSlug={leagueId}
        players={league.members ?? []}
        seasonId={activeSeason?.id}
      />
    </div>
  );
}
