import { api } from "@/lib/api";
import { CreateLeagueRequest, CreateLeagueResponse } from "./types";

export default async function createLeagueService(data: CreateLeagueRequest) {
  return await api
    .post("v1/leagues", {
      json: data,
    })
    .json<CreateLeagueResponse>();
}
