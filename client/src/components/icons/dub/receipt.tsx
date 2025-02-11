import { SVGProps, useEffect, useRef } from "react";

export function ReceiptIcon({
  "data-hovered": hovered,
  ...rest
}: { "data-hovered"?: boolean } & SVGProps<SVGSVGElement>) {
  const ref = useRef<SVGSVGElement>(null);

  useEffect(() => {
    if (!ref.current) return;

    if (hovered) {
      const lines = ref.current.querySelectorAll("line");

      lines.forEach((line, index) => {
        line.animate(
          [
            { opacity: 0, transform: "translateX(-5px)" },
            { opacity: 1, transform: "translateX(0px)" },
          ],
          {
            duration: 500, // Duração de 0.5 segundo para o "aparecimento"
            delay: index * 200, // Cada linha aparece com um atraso escalonado
            fill: "forwards", // Mantém a opacidade após a animação
            easing: "ease-out", // Suaviza o movimento final
          }
        );
      });
    } else {
      ref.current.getAnimations().forEach((anim) => anim.cancel()); // Cancela a animação quando o hover é removido
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
        <path d="M4 2h16v20H4z" /> {/* Bordas do recibo */}
        <path d="M4 2v20c0 0.5 0.5 1 1 1h14c0.5 0 1-0.5 1-1V2" />
        <line x1="8" y1="8" x2="16" y2="8" /> {/* Linha de texto 1 */}
        <line x1="8" y1="12" x2="16" y2="12" /> {/* Linha de texto 2 */}
        <line x1="8" y1="16" x2="14" y2="16" /> {/* Linha de texto 3 */}
      </g>
    </svg>
  );
}
