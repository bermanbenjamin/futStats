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
import { useEffect } from "react";
import useCreateLeagueForm from "./use-create-league-form";

export default function CreateLeagueForm() {
  const { player } = useSessionStore();
  const { form, onSubmit, isPending } = useCreateLeagueForm();

  useEffect(() => {
    form.setValue("ownerId", player!.id);
  }, [player, form]);

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <FormField
          control={form.control}
          name="name"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Nome da Liga</FormLabel>
              <FormControl>
                <Input {...field} />
              </FormControl>
            </FormItem>
          )}
        />
        <Button className="w-full mt-4" type="submit" disabled={isPending}>
          {isPending ? "Criando..." : "Criar"}
        </Button>
      </form>
    </Form>
  );
}
