import { SVGProps, useEffect, useRef } from "react";

export function CustomerIcon({
  "data-hovered": hovered,
  ...rest
}: { "data-hovered"?: boolean } & SVGProps<SVGSVGElement>) {
  const ref = useRef<SVGSVGElement>(null);

  useEffect(() => {
    if (!ref.current) return;

    const circles = ref.current.querySelectorAll("circle");

    if (hovered) {
      circles.forEach((circle, index) => {
        circle.animate([{ opacity: 0 }, { opacity: 1 }], {
          duration: 600,
          delay: index * 150, // Sequência de animação entre os círculos
          iterations: Infinity,
          direction: "alternate", // Vai e volta no efeito de fade
        });
      });
    } else {
      circles.forEach((circle) =>
        circle.getAnimations().forEach((anim) => anim.cancel())
      );
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
        {/* Cliente central */}
        <circle cx="12" cy="12" r="3" />
        {/* Clientes ao redor */}
        <circle cx="6" cy="12" r="2" />
        <circle cx="18" cy="12" r="2" />
        <circle cx="12" cy="6" r="2" />
        <circle cx="12" cy="18" r="2" />
      </g>
    </svg>
  );
}
