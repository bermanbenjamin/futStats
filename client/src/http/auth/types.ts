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

export type SignUpRequest = {
  name: string;
  email: string;
  password: string;
  age: number;
};

export type SignUpResponse = {
  token: string;
  user: Player;
};
