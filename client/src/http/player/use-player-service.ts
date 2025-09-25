import { useMutation, UseMutationOptions } from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import { toast } from "sonner";
import { getPlayerByEmail } from ".";
import { Player } from "../types";

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
