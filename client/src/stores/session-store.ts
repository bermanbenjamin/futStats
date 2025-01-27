import { Player } from "@/http/types";
import { create } from "zustand";

type SessionState = {
  player: Player | null;
  setPlayer: (player: Player) => void;
  removePlayer: () => void;
};

export const useSessionStore = create<SessionState>((set) => ({
  player: null,
  setPlayer: (player: Player) => set({ player }),
  removePlayer: () => set({ player: null }),
}));
