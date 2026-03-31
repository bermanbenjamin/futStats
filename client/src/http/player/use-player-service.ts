import { useMutation, UseMutationOptions, useQuery } from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import { toast } from "sonner";
import { getPlayerByEmail, getPlayerById } from ".";
import { Player } from "../types";

export function useGetPlayerById(id: string) {
  return useQuery({
    queryKey: ["player", id],
    queryFn: async (): Promise<Player> => {
      const res = await getPlayerById(id);
      return res.data;
    },
    enabled: !!id,
  });
}

export function useGetPlayerService(
  options?: UseMutationOptions<Player, HTTPError, string>
) {
  return useMutation({
    mutationKey: ["getPlayer"],
    mutationFn: getPlayerByEmail,
    onError: (error: Error) => {
      toast(error.message);
    },
    ...options,
  });
}
