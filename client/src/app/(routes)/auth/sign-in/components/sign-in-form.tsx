"use client";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import Link from "next/link";
import useSignInForm from "./use-sign-in-form";

export default function SignInForm() {
  const { form, onSubmit } = useSignInForm();

  return (
    <div className="flex flex-col gap-5">
      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <div className="w-full border-t" />
        </div>
      </div>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="flex flex-col space-y-4"
        >
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input placeholder="youremail@example.com" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <div className="flex justify-between">
                  <FormLabel>Senha</FormLabel>
                  <Link
                    href={"/auth/forgot-password"}
                    passHref
                    className="text-sm text-muted-foreground"
                  >
                    Esqueceu a senha?
                  </Link>
                </div>
                <div className="relative">
                  <FormControl>
                    <Input placeholder="••••••••" {...field} />
                  </FormControl>
                </div>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit">
            {/* {isPending && <Icons.loader className="animate-spin mr-2 size-4" />} */}
            Acessar
          </Button>
        </form>
      </Form>
    </div>
  );
}
