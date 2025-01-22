export type SignInResponse = {
  token: string;
  expires: string;
  user: object;
};

export type SignInRequest = {
  email: string;
  password: string;
};
