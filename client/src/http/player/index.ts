import { api } from "@/lib/api";
import { Player } from "../types";

async function getPlayerByEmail(email: string): Promise<Player> {
  const response = await api.get(`v1/players/${email}`, {
    headers: {
      "x-api-filter-type": "email",
    },
  });
  return response.json();
}

export { getPlayerByEmail };
