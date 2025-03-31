import reactLogo from "./assets/react.svg";
import goGopher from "./assets/go-gopher.svg";
import "./App.css";
import { useGetUsers, usePostUser } from "./hooks/useUsers";
import { Fragment, useState } from "react";
import UserRow from "./components/UserRow";

function App() {
  const [user, setUser] = useState({ id: 0, username: "" });

  const { data: users, isPending, error } = useGetUsers();
  const { mutate: postUser } = usePostUser();

  if (isPending) {
    return <div>loading...</div>;
  }

  if (error) {
    return (
      <div>
        <span>Error fetching lmao stuff</span>
        <span>{error.name}: {error.message}</span>
      </div>
    );
  }

  return (
    <>
      <div>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo" alt="React logo" />
        </a>
        <a href="https://go.dev/" target="_blank">
          <img src={goGopher} className="logo" alt="Go lang!" />
        </a>
      </div>
      <h1>React to deez + Go ligma</h1>
      <div className="add-user">
        <h2>
          Add user?
        </h2>
        <input
          onChange={(e) =>
            setUser({
              ...user,
              username: e.target.value,
            })}
          placeholder="lmao..."
        />
        <button onClick={() => postUser(user)}>Submit New User</button>
      </div>
      <br />
      <div className="card">
        <table>
          <thead>
            <tr>
              <th>Id</th>
              <th>User</th>
              <th>Update?</th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {users.map((u) => {
              return (
                <Fragment key={u.id}>
                  <UserRow u={u} />
                </Fragment>
              );
            })}
          </tbody>
        </table>
      </div>
    </>
  );
}

export default App;
