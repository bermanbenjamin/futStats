"use client";

import { FinishSeasonModal } from "@/components/modals/finish-season";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { getSeasonById } from "@/http/season/index";
import {
  useGetSeasonsByLeague,
  useGetSeasonStats,
} from "@/http/season/use-season-service";
import { useGetMatchesByLeague } from "@/http/match/use-match-service";
import { useSessionStore } from "@/stores/session-store";
import { useGetLeagueService } from "@/http/league/use-league-service";
import { useQuery } from "@tanstack/react-query";
import { ArrowLeft, Trophy, Star, Target } from "lucide-react";
import Link from "next/link";
import { use, useState } from "react";
import { Player } from "@/http/types";

export default function SeasonPage({
  params,
}: {
  params: Promise<{ playerId: string; leagueId: string; seasonId: string }>;
}) {
  const { playerId, leagueId, seasonId } = use(params);
  const [finishOpen, setFinishOpen] = useState(false);
  const { player: sessionPlayer } = useSessionStore();

  const { data: season, isLoading: loadingSeason } = useQuery({
    queryKey: ["season", seasonId],
    queryFn: () => getSeasonById(seasonId),
    enabled: !!seasonId,
  });

  const { data: stats, isLoading: loadingStats } = useGetSeasonStats(
    leagueId,
    seasonId
  );

  const { data: league } = useGetLeagueService(leagueId);
  const { data: allMatches } = useGetMatchesByLeague(leagueId);

  const isOwner = league?.ownerId === sessionPlayer?.id;

  // Filter matches for this season
  const seasonMatches = allMatches?.filter(
    (m) => (m as { season_id?: string }).season_id === seasonId
  );

  // Sort leaderboard by points (goals×3 + assists×1)
  const leaderboard = [...(stats ?? [])].sort((a, b) => {
    const pointsA = a.goals * 3 + a.assists;
    const pointsB = b.goals * 3 + b.assists;
    return pointsB - pointsA;
  });

  const topScorer = [...(stats ?? [])].sort((a, b) => b.goals - a.goals)[0];
  const topAssist = [...(stats ?? [])].sort(
    (a, b) => b.assists - a.assists
  )[0];

  if (loadingSeason) {
    return (
      <div className="w-full space-y-6 p-6">
        <Skeleton className="h-8 w-48" />
        <Skeleton className="h-24 w-full" />
        <Skeleton className="h-64 w-full" />
      </div>
    );
  }

  if (!season) {
    return (
      <div className="w-full p-6 text-center text-muted-foreground">
        Temporada não encontrada.
      </div>
    );
  }

  const formattedStart = season.init_date
    ? new Date(season.init_date).toLocaleDateString("pt-BR")
    : "—";
  const formattedEnd = season.end_date
    ? new Date(season.end_date).toLocaleDateString("pt-BR")
    : "—";

  return (
    <div className="w-full space-y-8 p-6">
      {/* Header */}
      <div className="space-y-3">
        <Link
          href={`/${playerId}/leagues/${leagueId}`}
          className="flex items-center gap-1 text-sm text-muted-foreground hover:text-foreground transition-colors"
        >
          <ArrowLeft className="w-4 h-4" />
          Voltar para a liga
        </Link>

        <div className="flex items-start justify-between">
          <div className="space-y-1">
            <div className="flex items-center gap-3">
              <h1 className="text-2xl font-bold">{season.year}</h1>
              {season.status === "active" ? (
                <Badge className="bg-green-500 hover:bg-green-600">Ativa</Badge>
              ) : (
                <Badge variant="secondary">Encerrada</Badge>
              )}
            </div>
            <p className="text-sm text-muted-foreground">
              {formattedStart} → {formattedEnd}
            </p>
          </div>

          {season.status === "active" && isOwner && (
            <Button
              variant="destructive"
              size="sm"
              onClick={() => setFinishOpen(true)}
            >
              Encerrar Temporada
            </Button>
          )}
        </div>
      </div>

      {/* Awards */}
      <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
        <AwardCard
          icon={<Trophy className="w-5 h-5 text-yellow-500" />}
          title="Artilheiro"
          player={topScorer}
          stat={topScorer ? `${topScorer.goals} gols` : undefined}
        />
        <AwardCard
          icon={<Star className="w-5 h-5 text-blue-500" />}
          title="Garçom"
          player={topAssist}
          stat={topAssist ? `${topAssist.assists} assistências` : undefined}
        />
        <AwardCard
          icon={<Target className="w-5 h-5 text-purple-500" />}
          title="Melhor em Campo"
          player={season.best_player ?? undefined}
          stat={undefined}
        />
      </div>

      {/* Leaderboard */}
      <div className="space-y-3">
        <h2 className="text-xl font-semibold">Classificação</h2>
        {loadingStats ? (
          <Skeleton className="h-48 w-full" />
        ) : leaderboard.length === 0 ? (
          <div className="text-center py-8 text-muted-foreground border rounded-lg">
            Nenhum evento registrado nesta temporada ainda.
          </div>
        ) : (
          <div className="overflow-x-auto rounded-lg border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-8">#</TableHead>
                  <TableHead>Jogador</TableHead>
                  <TableHead className="text-center">Gols</TableHead>
                  <TableHead className="text-center">Assist.</TableHead>
                  <TableHead className="text-center">Partidas</TableHead>
                  <TableHead className="text-center">C.A</TableHead>
                  <TableHead className="text-center">C.V</TableHead>
                  <TableHead className="text-center font-bold">Pts</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {leaderboard.map((player, i) => (
                  <TableRow key={player.id}>
                    <TableCell className="text-muted-foreground text-sm">
                      {i + 1}
                    </TableCell>
                    <TableCell className="font-medium">{player.name}</TableCell>
                    <TableCell className="text-center">{player.goals}</TableCell>
                    <TableCell className="text-center">{player.assists}</TableCell>
                    <TableCell className="text-center">{player.matches}</TableCell>
                    <TableCell className="text-center">{player.yellow_cards}</TableCell>
                    <TableCell className="text-center">{player.red_cards}</TableCell>
                    <TableCell className="text-center font-bold">
                      {player.goals * 3 + player.assists}
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>
        )}
      </div>

      {/* Matches */}
      <div className="space-y-3">
        <h2 className="text-xl font-semibold">
          Partidas ({seasonMatches?.length ?? 0})
        </h2>
        {!seasonMatches || seasonMatches.length === 0 ? (
          <div className="text-center py-6 text-muted-foreground border rounded-lg">
            Nenhuma partida registrada nesta temporada.
          </div>
        ) : (
          <div className="space-y-2">
            {seasonMatches.map((match) => (
              <div
                key={match.id}
                className="p-3 border rounded-lg flex justify-between items-center"
              >
                <span className="text-sm">
                  {new Date(match.date).toLocaleDateString("pt-BR", {
                    day: "2-digit",
                    month: "2-digit",
                    year: "numeric",
                  })}
                </span>
                <span className="text-sm text-muted-foreground">
                  {match.events?.length ?? 0} eventos
                </span>
              </div>
            ))}
          </div>
        )}
      </div>

      {season && (
        <FinishSeasonModal
          leagueSlug={leagueId}
          season={season}
          open={finishOpen}
          onOpenChange={setFinishOpen}
        />
      )}
    </div>
  );
}

function AwardCard({
  icon,
  title,
  player,
  stat,
}: {
  icon: React.ReactNode;
  title: string;
  player?: Player;
  stat?: string;
}) {
  return (
    <Card>
      <CardHeader className="pb-2">
        <CardTitle className="text-sm font-medium flex items-center gap-2 text-muted-foreground">
          {icon}
          {title}
        </CardTitle>
      </CardHeader>
      <CardContent>
        {player ? (
          <div>
            <p className="font-semibold">{player.name}</p>
            {stat && <p className="text-sm text-muted-foreground">{stat}</p>}
          </div>
        ) : (
          <p className="text-muted-foreground">—</p>
        )}
      </CardContent>
    </Card>
  );
}
