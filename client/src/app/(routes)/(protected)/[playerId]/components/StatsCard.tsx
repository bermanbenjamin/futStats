import type { LucideIcon } from "lucide-react";

function StatsCard({
  title,
  value,
  icon: IconComponent,
}: {
  title: string;
  value: number;
  icon: LucideIcon;
}) {
  return (
    <div className="flex gap-2 items-center bg-neutral-100 dark:bg-neutral-800 p-4 rounded-lg">
      <IconComponent className="w-6 h-6" />
      <h3 className="text-lg font-medium text-neutral-900 dark:text-neutral-100">
        {title} :
      </h3>
      <p className="text-2xl font-semibold text-indigo-900 dark:text-indigo-100">
        {value}
      </p>
    </div>
  );
}

export default StatsCard;
