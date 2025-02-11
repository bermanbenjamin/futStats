"use client";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useSessionStore } from "@/stores/session-store";
import useCreateLeagueForm from "./use-create-league-form";

export default function CreateLeagueForm() {
  const { player } = useSessionStore();
  const { form, onSubmit, isPending } = useCreateLeagueForm();

  form.setValue("ownerId", player!.ID);

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Nome</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
            </FormItem>
          )}
        />
        <Button type="submit" disabled={isPending}>
          {isPending ? "Criando..." : "Criar"}
        </Button>
      </form>
    </Form>
  );
}
