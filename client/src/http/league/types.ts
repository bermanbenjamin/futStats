import { League } from "../types";

export type CreateLeagueRequest = {
  ownerId: string;
  name: string;
};

export type CreateLeagueResponse = League;

export type GetLeagueResponse = League;
