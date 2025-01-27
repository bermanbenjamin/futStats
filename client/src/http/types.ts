export type ErrorResponse = {
  message: string;
};

export type Player = {
  CreatedAt: string; // ISO 8601 date string
  UpdatedAt: string; // ISO 8601 date string
  DeletedAt: string | null; // ISO 8601 date string or null
  ID: string; // UUID
  email: string;
  name: string;
  goals: number;
  assists: number;
  disarms: number;
  dribbles: number;
  matches: number;
  red_cards: number;
  yellow_cards: number;
  position: string;
};
