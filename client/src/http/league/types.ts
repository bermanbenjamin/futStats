import { League } from "../types";

export type CreateLeagueRequest = {
  ownerId: string;
  name: string;
};

export type CreateLeagueResponse = League;

export type GetLeagueResponse = League;

export type AddPlayerRequest = {
  email: string;
  slug: string;
};

// Mapping function for API responses
export function mapGetLeagueApiResponseToLeague(response: GetLeagueResponse): League {
  return response;
}
