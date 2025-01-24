import { useSignInService } from "@/http/auth/use-auth-service";
import { zodResolver } from "@hookform/resolvers/zod";
import { setCookie } from "cookies-next";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
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

  const { mutateAsync: signInService, isPending } = useSignInService({
    onSuccess: async (data) => {
      setCookie("token", data.token);
    },
    onError: (error) => {
      console.log(error);
      toast.error(error.message);
    },
  });

  async function onSubmit(data: SignInFormValues) {
    await signInService(data);
  }

  return { form, onSubmit, isPending };
}
