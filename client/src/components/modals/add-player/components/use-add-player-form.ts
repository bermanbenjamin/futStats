"use client";

import { useAddPlayerToLeague } from "@/http/league/use-league-service";
import { useGetPlayerService } from "@/http/player/use-player-service";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

const addPlayerFormSchema = z.object({
  email: z
    .string()
    .email("Invalid email address")
    .nonempty("Email is required"),
  slug: z.string(),
});

type AddPlayerFormData = z.infer<typeof addPlayerFormSchema>;

export default function useAddPlayerForm() {
  const form = useForm<AddPlayerFormData>({
    resolver: zodResolver(addPlayerFormSchema),
  });

  const { mutateAsync: addPlayerToLeague, isPending } = useAddPlayerToLeague();
  const { mutateAsync: getPlayer } = useGetPlayerService();

  async function onSubmit() {
    const data = form.getValues();
    const player = await getPlayer(data.email);

    if (!player) {
      toast("Jogador com este email não encontrado.");
      return;
    }

    await addPlayerToLeague(data);
  }

  return { form, onSubmit, isPending };
}
