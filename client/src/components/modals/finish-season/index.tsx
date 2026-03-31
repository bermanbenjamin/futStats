"use client";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Season } from "@/http/season/index";
import { useFinishSeason } from "@/http/season/use-season-service";
import { useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

interface Props {
  leagueSlug: string;
  season: Season;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function FinishSeasonModal({
  leagueSlug,
  season,
  open,
  onOpenChange,
}: Props) {
  const queryClient = useQueryClient();

  const { mutateAsync: finishSeason, isPending } = useFinishSeason({
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["seasons", leagueSlug] });
      toast.success("Temporada encerrada!");
      onOpenChange(false);
    },
  });

  async function handleConfirm() {
    await finishSeason({ leagueSlug, seasonId: season.id });
  }

  const formattedStart = season.init_date
    ? new Date(season.init_date).toLocaleDateString("pt-BR")
    : "—";
  const formattedEnd = season.end_date
    ? new Date(season.end_date).toLocaleDateString("pt-BR")
    : "—";

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[400px]">
        <DialogHeader>
          <DialogTitle>Encerrar {season.year}?</DialogTitle>
        </DialogHeader>
        <div className="space-y-4">
          <p className="text-sm text-muted-foreground">
            Esta ação encerrará a temporada{" "}
            <span className="font-medium">{season.year}</span> ({formattedStart}{" "}
            → {formattedEnd}). Isso não pode ser desfeito.
          </p>
          <div className="flex gap-3 justify-end">
            <Button
              variant="outline"
              onClick={() => onOpenChange(false)}
              disabled={isPending}
            >
              Cancelar
            </Button>
            <Button
              variant="destructive"
              onClick={handleConfirm}
              disabled={isPending}
            >
              {isPending ? "Encerrando..." : "Encerrar Temporada"}
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
