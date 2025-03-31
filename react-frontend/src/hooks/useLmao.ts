import { useQuery } from "@tanstack/react-query"
import { z } from "zod";

const lmao = z.object({
  title: z.string(),
});

type Lmao = z.infer<typeof lmao>;

async function getLmao(): Promise<Lmao> {
  const response = await fetch("http://localhost:8080/lmao");
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  const res = await response.json();
  const validatedRes = lmao.safeParse(res);

  if (!validatedRes.success) {
    console.error("Validation error:", validatedRes.error);

    throw new Error(`Shape of data from api did not match zod schema`);
  }

  return validatedRes.data;
}

export default function useLmao() {
  return useQuery({
    queryKey: ["lmao"],
    queryFn: getLmao,
    retry: false,
  })
}
