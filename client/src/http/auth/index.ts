import { api } from "@/lib/api";
import { SignInRequest, SignInResponse } from "./types";

export default async function signInService({
  email,
  password,
}: SignInRequest): Promise<SignInResponse> {
  return await api
    .post("v1/auth/login", {
      json: {
        username: email,
        password,
      },
    })
    .json();
}
