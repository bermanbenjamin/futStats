import {
  useMutation,
  UseMutationOptions,
  useQuery,
  UseQueryOptions,
} from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import { createLeagueService, getLeagueService } from ".";
import {
  CreateLeagueRequest,
  CreateLeagueResponse,
  GetLeagueResponse,
} from "./types";

function useCreateLeagueService(
  options?: UseMutationOptions<
    CreateLeagueResponse,
    HTTPError,
    CreateLeagueRequest
  >
) {
  return useMutation({
    mutationFn: async (data: CreateLeagueRequest) => {
      return createLeagueService(data);
    },
    mutationKey: ["createLeague"],
    ...options,
  });
}

function useGetLeagueService(
  id: string,
  options?: UseQueryOptions<GetLeagueResponse, HTTPError, string>
) {
  return useQuery({
    queryKey: ["league", id],
    queryFn: () => getLeagueService(id),
    ...options,
  });
}

export { useCreateLeagueService, useGetLeagueService };
