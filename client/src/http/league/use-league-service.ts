import {
  useMutation,
  UseMutationOptions,
  useQuery,
  UseQueryOptions,
} from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import { toast } from "sonner";
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
    onError: (error) => {
      console.error(error);
      toast.error(error.message);
    },
    ...options,
  });
}

function useGetLeagueService(
  id: string,
  options?: UseQueryOptions<GetLeagueResponse, HTTPError>
) {
  return useQuery({
    queryKey: ["league", id],
    queryFn: async () => await getLeagueService(id),
    ...options,
  });
}

export { useCreateLeagueService, useGetLeagueService };
