import {
  useMutation,
  UseMutationOptions,
  useQueryClient,
} from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import createLeagueService from ".";
import { CreateLeagueRequest, CreateLeagueResponse } from "./types";

export default function useCreateLeagueService(
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
