"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useSignInService } from "@/http/auth/use-auth-service";
import { useSessionStore } from "@/stores/session-store";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function SignInForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { setPlayer } = useSessionStore();
  const router = useRouter();

  const signInMutation = useSignInService({
    onSuccess: (data) => {
      console.log("Login successful, data:", data);
      console.log("Player ID:", data.player?.id);

      // Store token and player data
      localStorage.setItem("token", data.token);

      if (data.player?.id) {
        // Set player in session store to prevent sidebar redirect
        setPlayer(data.player);
        console.log("Player set in session store");

        console.log("Redirecting to:", `/${data.player.id}`);
        router.push(`/${data.player.id}`);
      } else {
        console.error("No player ID in response");
        alert("Login successful but no player data received");
      }
    },
    onError: (error) => {
      console.error("Login error:", error);
      alert("Login failed. Please try again.");
    },
  });

  const handleSubmit = async () => {
    console.log("Button clicked, handling login");

    // Basic form validation
    if (!email || !password) {
      alert("Please fill in all fields");
      return;
    }

    console.log("About to call mutation with:", { email, password });
    // Use the proper auth service
    signInMutation.mutate({ email, password });
    console.log("Mutation called");
  };

  return (
    <div className="flex flex-col gap-5">
      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <div className="w-full border-t" />
        </div>
      </div>
      <div className="flex flex-col space-y-4">
        <div>
          <label htmlFor="email" className="block text-sm font-medium mb-1">
            Email
          </label>
          <Input
            id="email"
            type="email"
            placeholder="youremail@example.com"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                handleSubmit();
              }
            }}
            required
          />
        </div>
        <div>
          <div className="flex justify-between mb-1">
            <label htmlFor="password" className="block text-sm font-medium">
              Senha
            </label>
            <span className="text-sm text-muted-foreground cursor-not-allowed opacity-50">
              Esqueceu a senha?
            </span>
          </div>
          <Input
            id="password"
            type="password"
            placeholder="••••••••"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                handleSubmit();
              }
            }}
            required
          />
        </div>
        <Button
          onClick={handleSubmit}
          disabled={signInMutation.isPending}
          className="w-full"
        >
          {signInMutation.isPending ? "Entrando..." : "Acessar"}
        </Button>
      </div>
    </div>
  );
}
