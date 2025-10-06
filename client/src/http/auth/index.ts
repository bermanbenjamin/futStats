import { api } from "@/lib/api";
import {
  SignInRequest,
  SignInResponse,
  SignUpRequest,
  SignUpResponse,
} from "./types";

export default async function signInService({
  email,
  password,
}: SignInRequest): Promise<SignInResponse> {
  return await api
    .post("/v1/auth/login", {
      json: {
        username: email,
        password,
      },
    })
    .json();
}

export async function signUpService({
  name,
  email,
  password,
  age,
}: SignUpRequest): Promise<SignUpResponse> {
  return await api
    .post("/v1/auth/signup", {
      json: {
        name,
        email,
        password,
        age,
      },
    })
    .json();
}
