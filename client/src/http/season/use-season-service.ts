import {
  useMutation,
  UseMutationOptions,
  useQuery,
} from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import { toast } from "sonner";
import {
  createSeason,
  finishSeason,
  getSeasonsByLeague,
  getSeasonStats,
  Season,
  CreateSeasonRequest,
} from ".";
import { Player } from "@/http/types";

const handleApiError = (error: HTTPError): void => {
  console.error("API Error:", {
    status: error.response?.status,
    message: error.message,
  });
  toast.error(error.message || "An unexpected error occurred.");
};

function useGetSeasonsByLeague(leagueSlug: string) {
  return useQuery({
    queryKey: ["seasons", leagueSlug],
    queryFn: () => getSeasonsByLeague(leagueSlug),
    enabled: !!leagueSlug,
  });
}

function useCreateSeason(
  options?: UseMutationOptions<
    Season,
    HTTPError,
    { leagueSlug: string; data: CreateSeasonRequest }
  >
) {
  return useMutation({
    mutationFn: ({ leagueSlug, data }: { leagueSlug: string; data: CreateSeasonRequest }) =>
      createSeason(leagueSlug, data),
    mutationKey: ["createSeason"],
    onError: handleApiError,
    ...options,
  });
}

function useGetSeasonStats(leagueSlug: string, seasonId: string) {
  return useQuery<Player[], HTTPError>({
    queryKey: ["season-stats", leagueSlug, seasonId],
    queryFn: () => getSeasonStats(leagueSlug, seasonId),
    enabled: !!leagueSlug && !!seasonId,
  });
}

function useFinishSeason(
  options?: UseMutationOptions<
    Season,
    HTTPError,
    { leagueSlug: string; seasonId: string }
  >
) {
  return useMutation({
    mutationFn: ({ leagueSlug, seasonId }: { leagueSlug: string; seasonId: string }) =>
      finishSeason(leagueSlug, seasonId),
    mutationKey: ["finishSeason"],
    onError: handleApiError,
    ...options,
  });
}

export {
  useGetSeasonsByLeague,
  useCreateSeason,
  useGetSeasonStats,
  useFinishSeason,
};
