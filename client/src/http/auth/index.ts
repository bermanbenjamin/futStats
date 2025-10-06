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
  console.log("signInService called with:", { email, password });
  console.log("API base URL:", process.env.NEXT_PUBLIC_API_URL);

  const response = (await api
    .post("/v1/auth/login", {
      json: {
        username: email,
        password,
      },
    })
    .json()) as SignInResponse;

  console.log("Login response:", response);

  return response;
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
