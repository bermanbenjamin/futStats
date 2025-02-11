import { useMutation, UseMutationOptions } from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import signInService from ".";
import { ErrorResponse } from "../types";
import { SignInRequest, SignInResponse } from "./types";

export function useSignInService(
  options?: UseMutationOptions<
    SignInResponse,
    HTTPError<ErrorResponse>,
    SignInRequest
  >
) {
  return useMutation({
    mutationKey: ["signIn"],
    mutationFn: async (data: SignInRequest) => {
      return signInService(data);
    },
    ...options,
  });
}
