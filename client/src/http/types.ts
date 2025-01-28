export type ErrorResponse = {
  message: string;
};

type Base = {
  ID: string; // UUID
  CreatedAt: string; // ISO 8601 date string
  UpdatedAt: string; // ISO 8601 date string
  DeletedAt: string | null; // ISO 8601 date string or null
};

export type Player = {
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
  leagues: League[] | null;
};

type Season = {
  ID: string; // UUID
  name: string;
  startDate: string; // ISO 8601 date string
  endDate: string; // ISO 8601 date string
  leagues: League[]; // Many-to-many relationship with League
};

type League = Base & {
  ownerId: string; // UUID (references Player.ID)
  owner: Player; // Foreign key relationship
  name: string;
  slug: string;
  seasons: Season[]; // Many-to-many relationship with Season
  players: Player[]; // Many-to-many relationship with Player
};
