import { SVGProps, useEffect, useRef } from "react";

export function LayersIcon({
  "data-hovered": hovered,
  ...rest
}: { "data-hovered"?: boolean } & SVGProps<SVGSVGElement>) {
  const ref = useRef<SVGSVGElement>(null);

  useEffect(() => {
    if (!ref.current) return;

    if (hovered) {
      const layers = ref.current.querySelectorAll("polygon, polyline");

      layers.forEach((layer, index) => {
        layer.animate(
          [
            { transform: `translateY(0px)` },
            { transform: `translateY(-2px)` },
            { transform: `translateY(0px)` },
          ],
          {
            duration: 1200 + index * 150, // Duração mais longa para movimento suave
            easing: "ease-in-out",
          }
        );
      });
    } else {
      ref.current.getAnimations().forEach((anim) => anim.cancel()); // Cancela a animação ao sair do hover
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
        <polygon points="12 2 2 7 12 12 22 7 12 2" />
        <polyline points="2 17 12 22 22 17" />
        <polyline points="2 12 12 17 22 12" />
      </g>
    </svg>
  );
}
