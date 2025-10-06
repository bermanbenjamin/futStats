import { useMutation, UseMutationOptions } from "@tanstack/react-query";
import { HTTPError } from "ky-universal";
import signInService, { signUpService } from ".";
import { ErrorResponse } from "../types";
import {
  SignInRequest,
  SignInResponse,
  SignUpRequest,
  SignUpResponse,
} from "./types";

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

export function useSignUpService(
  options?: UseMutationOptions<
    SignUpResponse,
    HTTPError<ErrorResponse>,
    SignUpRequest
  >
) {
  return useMutation({
    mutationKey: ["signUp"],
    mutationFn: async (data: SignUpRequest) => {
      return signUpService(data);
    },
    ...options,
  });
}
