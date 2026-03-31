import { api } from "@/lib/api";
import { Player } from "../types";

async function getPlayerByEmail(email: string): Promise<Player> {
  const response = await api.get(`v1/players/${email}`, {
    headers: {
      "X-api-field-type": "email",
    },
  });
  return response.json();
}

async function getPlayerById(id: string): Promise<{ data: Player }> {
  return api.get(`v1/players/${id}`).json<{ data: Player }>();
}

export { getPlayerByEmail, getPlayerById };
