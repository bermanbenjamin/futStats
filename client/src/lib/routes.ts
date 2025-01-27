export const appRoutes = {
  home: "/",
  auth: {
    signIn: "/auth/sign-in",
    register: "/register",
    forgotPassword: "/forgot-password",
    resetPassword: "/reset-password",
  },
  player: {
    home: (id: string) => `/${id}`,
    search: "/player/search",
    library: "/player/library",
    settings: "/player/settings",
  },
};
