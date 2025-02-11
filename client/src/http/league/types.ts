import { League } from "../types";

export type CreateLeagueRequest = {
  ownerId: string;
  name: string;
};

export type CreateLeagueResponse = League;
