import { api } from "@/lib/api";
import { SignInRequest, SignInResponse } from "./types";

export default async function signInService({
  email,
  password,
}: SignInRequest): Promise<SignInResponse> {
  return await api
    .post("api/auth/sign-in", {
      json: {
        email,
        password,
      },
    })
    .json<SignInResponse>();
}
