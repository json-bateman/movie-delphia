import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { z } from "zod";

const user = z.object({
  id: z.number(),
  username: z.string(),
});
const users = z.array(user);

export type User = z.infer<typeof user>;
export type Users = User[];

async function getUsers(): Promise<Users> {
  const response = await fetch("http://localhost:8080/users");
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  const res = await response.json();
  const validatedRes = users.safeParse(res);

  if (!validatedRes.success) {
    console.error("Validation error:", validatedRes.error);

    throw new Error(`Shape of data from api did not match zod schema`);
  }

  return validatedRes.data;
}

async function postUser(user: User) {
  const response = await fetch(`http://localhost:8080/users`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json", // Inform the server the payload is JSON
    },
    body: JSON.stringify(user),
  });
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
}

async function putUser(user: User) {
  const response = await fetch(`http://localhost:8080/users/${user.id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json", // Inform the server the payload is JSON
    },
    body: JSON.stringify(user),
  });
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
}

async function deleteUser(id: number) {
  const response = await fetch(`http://localhost:8080/users/${id}`, {
    method: "DELETE",
  });
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
}

/* Hooks */
export function useGetUsers() {
  return useQuery({
    queryKey: ["users"],
    queryFn: getUsers,
    retry: false,
  });
}

export function usePostUser() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (user: User) => postUser(user),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["users"],
      });
    },
  });
}

export function usePutUser() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (user: User) => putUser(user),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["users"],
      });
    },
  });
}

export function useDeleteUser() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => deleteUser(id),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["users"],
      });
    },
  });
}
