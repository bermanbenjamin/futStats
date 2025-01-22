import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

const signInFormSchema = z.object({
  email: z.string({ required_error: "Email é obrigatório" }).email({
    message: "Email inválido",
  }),
  password: z.string({ required_error: "Senha é obrigatória" }).min(6, {
    message: "Senha deve ter ao menos 6 caracteres",
  }),
});

type SignInFormValues = z.infer<typeof signInFormSchema>;

export default function useSignInForm() {
  const form = useForm<SignInFormValues>({
    resolver: zodResolver(signInFormSchema),
  });

  async function onSubmit(data: SignInFormValues) {
    console.log(data);
  }

  return { form, onSubmit };
}
