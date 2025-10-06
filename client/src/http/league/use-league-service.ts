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

// Pure error handling function with explicit side effects
const handleApiError = (error: HTTPError): void => {
  // Explicit side effect: console logging
  console.error("API Error:", {
    status: error.response?.status,
    message: error.message,
    url: error.response?.url,
  });

  // Explicit side effect: user notification
  const userMessage =
    error.response?.status === 401
      ? "Authentication required. Please sign in."
      : error.response?.status === 403
      ? "You don't have permission to perform this action."
      : error.response?.status === 404
      ? "Resource not found."
      : error.response?.status >= 500
      ? "Server error. Please try again later."
      : error.message || "An unexpected error occurred.";

  toast.error(userMessage);
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
    onError: handleApiError,
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
    onError: handleApiError,
    ...options,
  });
}

export { useAddPlayerToLeague, useCreateLeagueService, useGetLeagueService };
