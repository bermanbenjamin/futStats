import { api } from "@/lib/api";
import { Player } from "@/http/types";

export type Season = {
  id: string;
  created_at: string;
  updated_at: string;
  year: string;
  init_date: string;
  end_date: string;
  status: "active" | "finished";
  goals_amount: number;
  assists_amount: number;
  striker_id: string | null;
  striker: Player | null;
  waiter_id: string | null;
  waiter: Player | null;
  best_player_id: string | null;
  best_player: Player | null;
};

export type CreateSeasonRequest = {
  name: string;
  init_date: string;
  end_date: string;
};

async function createSeason(
  leagueSlug: string,
  data: CreateSeasonRequest
): Promise<Season> {
  return api
    .post(`v1/leagues/${leagueSlug}/seasons`, {
      json: {
        year: data.name,
        init_date: data.init_date,
        end_date: data.end_date,
      },
    })
    .json<Season>();
}

async function getSeasonsByLeague(leagueSlug: string): Promise<Season[]> {
  return api.get(`v1/leagues/${leagueSlug}/seasons`).json<Season[]>();
}

async function getSeasonById(seasonId: string): Promise<Season> {
  return api.get(`v1/seasons/${seasonId}`).json<Season>();
}

async function getSeasonStats(
  leagueSlug: string,
  seasonId: string
): Promise<Player[]> {
  return api
    .get(`v1/leagues/${leagueSlug}/seasons/${seasonId}/stats`)
    .json<Player[]>();
}

async function finishSeason(
  leagueSlug: string,
  seasonId: string
): Promise<Season> {
  return api
    .post(`v1/leagues/${leagueSlug}/seasons/${seasonId}/finish`)
    .json<Season>();
}

export {
  createSeason,
  getSeasonsByLeague,
  getSeasonById,
  getSeasonStats,
  finishSeason,
};
