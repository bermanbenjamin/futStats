import { Player } from "../types";

export type SignInResponse = {
  token: string;
  expires: string;
  player: Player;
};

export type SignInRequest = {
  email: string;
  password: string;
};
