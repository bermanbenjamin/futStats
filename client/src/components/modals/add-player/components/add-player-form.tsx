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
import { useEffect } from "react";
import useAddPlayerForm from "./use-add-player-form";

interface AddPlayerFormProps {
  slug: string;
}

export default function AddPlayerForm({ slug }: AddPlayerFormProps) {
  const { form, isPending, onSubmit } = useAddPlayerForm();

  useEffect(() => {
    form.setValue("slug", slug);
  }, [slug, form]);

  return (
    <Form {...form}>
      <form onSubmit={onSubmit}>
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email do Jogador</FormLabel>
              <FormControl>
                <Input placeholder="youremail@example.com" {...field} />
              </FormControl>
            </FormItem>
          )}
        />
        <Button className="w-full mt-4" type="submit" disabled={isPending}>
          Adicionar Jogador
        </Button>
      </form>
    </Form>
  );
}
