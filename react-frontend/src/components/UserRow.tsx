import { useState } from "react";
import { useDeleteUser, usePutUser, User } from "../hooks/useUsers";

export default function UserRow({ u }: { u: User }) {
  const [user, setUser] = useState(u);
  const { mutate: deleteUser } = useDeleteUser();
  const { mutate: putUser } = usePutUser();

  return (
    <tr>
      <td>
        <span>{u.id}</span>
      </td>
      <td>
        <input
          value={user.username}
          onChange={(e) =>
            setUser({
              ...user,
              username: e.target.value,
            })}
        />
      </td>
      <td>
        <button onClick={() => putUser(user)}>Update User</button>
      </td>
      <td>
        <span
          style={{ color: "red", cursor: "pointer" }}
          onClick={() => deleteUser(u.id)}
        >
          X
        </span>
      </td>
    </tr>
  );
}
