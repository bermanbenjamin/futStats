import { api } from "@/lib/api";

export type Match = {
  id: string;
  league_id: string;
  season_id?: string;
  date: string;
  createdAt: string;
  events: MatchEvent[];
};

export type MatchEvent = {
  id: string;
  type: string;
  player_id: string;
  match_id: string;
  assistent_id?: string;
};

export type CreateMatchRequest = {
  league_id: string;
  date: string;
  season_id?: string;
};

export type CreateEventRequest = {
  type: string;
  player_id: string;
  match_id: string;
  assistent_id?: string;
};

async function createMatch(data: CreateMatchRequest): Promise<{ data: Match }> {
  return api
    .post(`v1/leagues/${data.league_id}/matches`, {
      json: {
        league_id: data.league_id,
        date: data.date,
        ...(data.season_id ? { season_id: data.season_id } : {}),
      },
    })
    .json<{ data: Match }>();
}

async function getMatchesByLeague(
  leagueSlug: string
): Promise<{ data: Match[] }> {
  return api.get(`v1/leagues/${leagueSlug}/matches`).json<{ data: Match[] }>();
}

async function createEvent(
  data: CreateEventRequest
): Promise<{ data: MatchEvent }> {
  return api.post("v1/events", { json: data }).json<{ data: MatchEvent }>();
}

export { createEvent, createMatch, getMatchesByLeague };
