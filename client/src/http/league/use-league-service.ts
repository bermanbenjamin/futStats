import {
  useMutation,
  UseMutationOptions,
  useQuery,
  useQueryClient,
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
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async (data: CreateLeagueRequest) => {
      return createLeagueService(data);
    },
    mutationKey: ["createLeague"],
    ...options,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["leagues"] });
    },
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
