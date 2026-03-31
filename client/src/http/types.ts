export type ErrorResponse = {
  message: string;
};

type Base = {
  id: string; // UUID
  createdAt: string; // ISO 8601 date string
  updatedAt: string; // ISO 8601 date string
  deletedAt: string | null; // ISO 8601 date string or null
};

export type Player = {
  id: string; // UUID
  email: string;
  name: string;
  position: string;
  goals: number;
  assists: number;
  disarms: number;
  dribbles: number;
  matches: number;
  red_cards: number;
  yellow_cards: number;
  member_of_leagues: League[] | null;
  owned_leagues: League[] | null;
};

export type Season = {
  id: string;
  created_at: string;
  updated_at: string;
  year: string;
  init_date: string;
  end_date: string;
  status: "active" | "finished";
  goals_amount: number;
  assists_amount: number;
  striker_id: string | null;
  striker: Player | null;
  waiter_id: string | null;
  waiter: Player | null;
  best_player_id: string | null;
  best_player: Player | null;
};

export type League = Base & {
  ownerId: string; // UUID (references Player.ID)
  owner: Player; // Foreign key relationship
  name: string;
  slug: string;
  seasons: Season[] | null;
  members: Player[]; // Many-to-many relationship with Player
};
