export type ErrorResponse = {
  message: string;
};

type Base = {
  id: string; // UUID
  CreatedAt: string; // ISO 8601 date string
  UpdatedAt: string; // ISO 8601 date string
  DeletedAt: string | null; // ISO 8601 date string or null
};

export type Player = {
  id: string; // UUID
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
  member_leagues: League[] | null;
  owned_leagues: League[] | null;
};

type Season = {
  ID: string; // UUID
  name: string;
  startDate: string; // ISO 8601 date string
  endDate: string; // ISO 8601 date string
  leagues: League[]; // Many-to-many relationship with League
};

export type League = Base & {
  ownerId: string; // UUID (references Player.ID)
  owner: Player; // Foreign key relationship
  name: string;
  slug: string;
  seasons: Season[]; // Many-to-many relationship with Season
  members: Player[]; // Many-to-many relationship with Player
};
