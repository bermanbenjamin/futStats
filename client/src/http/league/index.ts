import { api } from "@/lib/api";
import {
  CreateLeagueRequest,
  CreateLeagueResponse,
  GetLeagueResponse,
} from "./types";

async function createLeagueService(data: CreateLeagueRequest) {
  return await api
    .post("v1/leagues", {
      json: {
        owner_id: data.ownerId,
        name: data.name,
      },
    })
    .json<CreateLeagueResponse>();
}

async function getLeagueService(id: string) {
  return await api.get(`v1/leagues/${id}`).json<GetLeagueResponse>();
}

async function addPlayerToLeague(email: string, slug: string) {
  return await api
    .post(`v1/${slug}/players`, {
      json: {
        email,
      },
    })
    .json<GetLeagueResponse>();
}

export { addPlayerToLeague, createLeagueService, getLeagueService };
