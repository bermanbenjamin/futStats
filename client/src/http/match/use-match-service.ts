import {
  useMutation,
  UseMutationOptions,
  useQuery,
} from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import { toast } from "sonner";
import {
  createEvent,
  createMatch,
  CreateEventRequest,
  CreateMatchRequest,
  getMatchesByLeague,
  Match,
  MatchEvent,
} from ".";

function useCreateMatch(
  options?: UseMutationOptions<{ data: Match }, HTTPError, CreateMatchRequest>
) {
  return useMutation({
    mutationFn: createMatch,
    mutationKey: ["createMatch"],
    onError: () => toast.error("Failed to create match"),
    ...options,
  });
}

function useGetMatchesByLeague(leagueSlug: string) {
  return useQuery({
    queryKey: ["matches", leagueSlug],
    queryFn: () => getMatchesByLeague(leagueSlug),
    enabled: !!leagueSlug,
    select: (res) => res.data,
  });
}

function useCreateEvent(
  options?: UseMutationOptions<
    { data: MatchEvent },
    HTTPError,
    CreateEventRequest
  >
) {
  return useMutation({
    mutationFn: createEvent,
    mutationKey: ["createEvent"],
    onError: () => toast.error("Failed to record event"),
    ...options,
  });
}

export { useCreateEvent, useCreateMatch, useGetMatchesByLeague };
