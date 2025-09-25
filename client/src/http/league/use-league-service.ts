import {
  useMutation,
  UseMutationOptions,
  useQuery,
  UseQueryOptions,
} from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import { toast } from "sonner";
import { addPlayerToLeague, createLeagueService, getLeagueService } from ".";
import {
  AddPlayerRequest,
  CreateLeagueRequest,
  CreateLeagueResponse,
  GetLeagueResponse,
  mapGetLeagueApiResponseToLeague,
} from "./types";

// Error handling function
const handleError = (error: HTTPError) => {
  console.error(error);
  toast.error(error.message);
};

function useCreateLeagueService(
  options?: UseMutationOptions<
    CreateLeagueResponse,
    HTTPError,
    CreateLeagueRequest
  >
) {
  return useMutation({
    mutationFn: createLeagueService,
    mutationKey: ["createLeague"],
    onError: handleError,
    ...options,
  });
}

function useGetLeagueService(
  id: string,
  options?: UseQueryOptions<GetLeagueResponse, HTTPError>
) {
  return useQuery({
    queryKey: ["league", id],
    queryFn: async () => {
      const response = await getLeagueService(id);
      return mapGetLeagueApiResponseToLeague(response);
    },
    ...options,
  });
}

function useAddPlayerToLeague(
  options?: UseMutationOptions<GetLeagueResponse, HTTPError, AddPlayerRequest>
) {
  return useMutation({
    mutationFn: (data: AddPlayerRequest) =>
      addPlayerToLeague(data.email, data.slug),
    mutationKey: ["leagues"],
    onError: handleError,
    ...options,
  });
}

export { useAddPlayerToLeague, useCreateLeagueService, useGetLeagueService };
