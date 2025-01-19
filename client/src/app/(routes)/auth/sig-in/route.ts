import { api } from "@/lib/api";

type SignInResponse = {
  token: string;
};

export default async function signIn(email: string, password: string) {
  const response = await api.post("auth/sign-in", {
    json: { email, password },
  });
  const data = await response.json<SignInResponse>();
  setupToken(data);
  return data;
}

function setupToken(data: SignInResponse) {
  localStorage.setItem("token", data.token);
}
