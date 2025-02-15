import { useCreateLeagueService } from "@/http/league/use-league-service";
import { League } from "@/http/types";
import { appRoutes } from "@/lib/routes";
import { useSessionStore } from "@/stores/session-store";
import { zodResolver } from "@hookform/resolvers/zod";
import { useQueryClient } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { z } from "zod";

const createLeagueFormSchema = z.object({
  ownerId: z.string().min(1, { message: "ID do proprietário é obrigatório" }),
  name: z.string().min(1, { message: "Nome é obrigatório" }),
});

export default function useCreateLeagueForm() {
  const form = useForm<z.infer<typeof createLeagueFormSchema>>({
    resolver: zodResolver(createLeagueFormSchema),
  });

  const { player, setPlayer } = useSessionStore();
  const router = useRouter();
  const queryClient = useQueryClient();

  const { mutateAsync: createLeagueService, isPending } =
    useCreateLeagueService({
      onSuccess: (league: League) => {
        setPlayer({
          ...player!,
          owned_leagues: player!.owned_leagues!.concat(league),
        });
        router.push(appRoutes.player.leagues.get(league.id));
        queryClient.invalidateQueries({ queryKey: ["leagues"] });
      },
    });

  async function onSubmit(data: z.infer<typeof createLeagueFormSchema>) {
    await createLeagueService(data);
  }

  return { form, onSubmit, isPending };
}
