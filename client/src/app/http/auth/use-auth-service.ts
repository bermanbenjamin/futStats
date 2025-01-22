import { useMutation, UseMutationOptions } from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import signInService from ".";
import { SignInRequest, SignInResponse } from "./types";

export function useSignInService(
  options?: UseMutationOptions<SignInResponse, HTTPError<object>, SignInRequest>
) {
  return useMutation({
    mutationKey: ["signIn"],
    mutationFn: async (data: SignInRequest) => {
      return signInService(data);
    },
    ...options,
  });
}
