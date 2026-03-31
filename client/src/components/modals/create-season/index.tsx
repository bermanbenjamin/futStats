"use client";

import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useCreateSeason } from "@/http/season/use-season-service";
import { zodResolver } from "@hookform/resolvers/zod";
import { useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { useForm } from "react-hook-form";
import { z } from "zod";

const schema = z
  .object({
    name: z.string().min(1, "Nome é obrigatório"),
    init_date: z.string().min(1, "Data de início é obrigatória"),
    end_date: z.string().min(1, "Data de fim é obrigatória"),
  })
  .refine((d) => d.end_date > d.init_date, {
    message: "Data de fim deve ser após a data de início",
    path: ["end_date"],
  });

type FormValues = z.infer<typeof schema>;

interface Props {
  leagueSlug: string;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export function CreateSeasonModal({ leagueSlug, open, onOpenChange }: Props) {
  const queryClient = useQueryClient();

  const form = useForm<FormValues>({
    resolver: zodResolver(schema),
    defaultValues: { name: "", init_date: "", end_date: "" },
  });

  const { mutateAsync: createSeason, isPending } = useCreateSeason({
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["seasons", leagueSlug] });
      toast.success("Temporada criada com sucesso!");
      form.reset();
      onOpenChange(false);
    },
  });

  async function onSubmit(values: FormValues) {
    await createSeason({ leagueSlug, data: values });
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Iniciar Temporada</DialogTitle>
        </DialogHeader>
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Nome da Temporada</FormLabel>
                  <FormControl>
                    <Input placeholder="Ex: Temporada 2025" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="init_date"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Data de Início</FormLabel>
                  <FormControl>
                    <Input type="date" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="end_date"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Data de Fim</FormLabel>
                  <FormControl>
                    <Input type="date" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit" className="w-full" disabled={isPending}>
              {isPending ? "Criando..." : "Criar Temporada"}
            </Button>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
