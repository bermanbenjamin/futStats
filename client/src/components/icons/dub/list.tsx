import { SVGProps, useEffect, useRef } from "react";

export function ListIcon({
  "data-hovered": hovered,
  ...rest
}: { "data-hovered"?: boolean } & SVGProps<SVGSVGElement>) {
  const ref = useRef<SVGSVGElement>(null);

  useEffect(() => {
    if (!ref.current) return;

    if (hovered) {
      // Animação de deslizamento para cada linha
      const lines = ref.current.querySelectorAll("line");
      lines.forEach((line, index) => {
        line.animate(
          [
            { transform: "translateX(0)" },
            { transform: `translateX(${index % 2 === 0 ? 10 : -10}px)` },
            { transform: "translateX(0)" },
          ],
          {
            duration: 600, // Duração de 0.6s por ciclo
            delay: index * 100, // Pequeno atraso entre cada linha
            iterations: Infinity, // Repetição contínua
          }
        );
      });
    } else {
      // Cancela todas as animações em execução
      ref.current
        .querySelectorAll("line")
        .forEach((line) => line.getAnimations().forEach((anim) => anim.cancel()));
    }
  }, [hovered]);

  return (
    <svg
      height="24"
      width="24"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
      ref={ref}
      {...rest}
    >
      <g
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <line x1="8" y1="6" x2="21" y2="6" />
        <line x1="8" y1="12" x2="21" y2="12" />
        <line x1="8" y1="18" x2="21" y2="18" />
        <circle cx="4" cy="6" r="1" />
        <circle cx="4" cy="12" r="1" />
        <circle cx="4" cy="18" r="1" />
      </g>
    </svg>
  );
}
