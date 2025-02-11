import { useCreateLeagueService } from "@/http/league/use-league-service";
import { League } from "@/http/types";
import { appRoutes } from "@/lib/routes";
import { zodResolver } from "@hookform/resolvers/zod";
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

  const router = useRouter();

  const { mutateAsync: createLeagueService, isPending } =
    useCreateLeagueService({
      onSuccess: (league: League) => {
        router.push(appRoutes.player.leagues.get(league.id));
      },
    });

  async function onSubmit(data: z.infer<typeof createLeagueFormSchema>) {
    console.log(data);
    await createLeagueService(data);
  }

  return { form, onSubmit, isPending };
}
