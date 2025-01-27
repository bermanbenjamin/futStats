import { SVGProps, useEffect, useRef } from "react";

export function HomeIcon({
  "data-hovered": hovered,
  ...rest
}: { "data-hovered"?: boolean } & SVGProps<SVGSVGElement>) {
  const ref = useRef<SVGSVGElement>(null);

  useEffect(() => {
    if (!ref.current) return;

    if (hovered) {
      ref.current.animate(
        [{ transform: "rotate(0deg)" }, { transform: "rotate(360deg)" }],
        {
          duration: 1000, // Duração de 1 segundo para o giro completo
          iterations: Infinity, // Gira continuamente
        }
      );
    } else {
      ref.current.getAnimations().forEach((anim) => anim.cancel()); // Para a animação quando o hover é removido
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
        <circle cx="12" cy="12" r="10" />
        <path d="M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20" />
        <path d="M2 12h20" />
      </g>
    </svg>
  );
}
