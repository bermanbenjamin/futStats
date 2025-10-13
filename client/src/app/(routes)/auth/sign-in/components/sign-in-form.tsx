"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useSignInService } from "@/http/auth/use-auth-service";
import Link from "next/link";
import { useState } from "react";

export default function SignInForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const signInMutation = useSignInService({
    onSuccess: (data) => {
      // Store token and redirect to player dashboard
      localStorage.setItem("token", data.token);
      window.location.href = `/${data.player.id}`;
    },
    onError: (error) => {
      console.error("Login error:", error);
      alert("Login failed. Please try again.");
    },
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Basic form validation
    if (!email || !password) {
      alert("Please fill in all fields");
      return;
    }

    // Use the proper auth service
    signInMutation.mutate({ email, password });
  };

  return (
    <div className="flex flex-col gap-5">
      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <div className="w-full border-t" />
        </div>
      </div>
      <form onSubmit={handleSubmit} className="flex flex-col space-y-4">
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
            required
          />
        </div>
        <div>
          <div className="flex justify-between mb-1">
            <label htmlFor="password" className="block text-sm font-medium">
              Senha
            </label>
            <Link
              href="/auth/forgot-password"
              className="text-sm text-muted-foreground hover:underline"
            >
              Esqueceu a senha?
            </Link>
          </div>
          <Input
            id="password"
            type="password"
            placeholder="••••••••"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <Button
          type="submit"
          disabled={signInMutation.isPending}
          className="w-full"
        >
          {signInMutation.isPending ? "Entrando..." : "Acessar"}
        </Button>
      </form>
    </div>
  );
}
