"use client";

import { Icons } from "@/components/icons";
import { Input } from "@/components/ui/input";
import { useState } from "react";

type PasswordInputProps = React.ComponentProps<typeof Input>;

export default function PasswordInput({ ...props }: PasswordInputProps) {
  const [isVisible, setIsVisible] = useState(false);

  function toggleVisibility() {
    setIsVisible(!isVisible);
  }

  return (
    <div className="relative">
      <Input
        {...props}
        className="pr-9"
        type={isVisible ? "text" : "password"}
      />
      <button
        className="absolute inset-y-px right-px flex h-full w-9 items-center justify-center rounded-r-lg text-muted-foreground/80 ring-offset-background transition-shadow hover:text-foreground focus-visible:border focus-visible:border-ring focus-visible:text-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring/30 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
        type="button"
        onClick={() => toggleVisibility()}
        aria-label={isVisible ? "Ocultar senha" : "Exibir senha"}
        aria-pressed={isVisible}
        aria-controls="password"
      >
        {isVisible ? (
          <Icons.eye
            className="size-4"
            strokeWidth={2}
            aria-hidden="true"
            role="presentation"
          />
        ) : (
          <Icons.eyeOff
            className="size-4"
            strokeWidth={2}
            aria-hidden="true"
            role="presentation"
          />
        )}
      </button>
    </div>
  );
}
