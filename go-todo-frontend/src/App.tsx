import { useEffect, useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import { fetchTodos, fetchMe, fetchPokemon } from "./api";
import { Todo, Post, Pokemon } from "./types";

import ElementList from "./components/ElementList";

function App() {
  const [count, setCount] = useState<number>(0);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [posts, setPosts] = useState<Post[]>([]);
  const [pokemon, setPokemon] = useState<Pokemon>({
    name: "Unknown",
    height: 0,
    weight: 0,
  });
  useEffect(() => {
    fetchTodos().then(setTodos);
  }, []);

  useEffect(() => {
    fetchMe().then(setPosts);
  }, []);

  useEffect(() => {
    fetchPokemon().then(setPokemon);
  }, []);

  const todo_headers = todos.length > 0 ? Object.keys(todos[0]) : [];
  const post_headers = posts.length > 0 ? Object.keys(posts[0]) : [];
  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p>
      </div>
      <div className="p-4">
        {/* <h1 className="text-xl font-bold mb-4">Todo List</h1> */}
        {/* <ul>
          {todos.map((todo) => (
            <li key={todo.ID} className="p-2 border-b">
              {todo.Title} {todo.Completed ? "✅" : "❌"}
            </li>
          ))}
        </ul> */}
        <ElementList e_name="Todo" elements={todos} headers={todo_headers} />
        <ElementList e_name="Posts" elements={posts.slice(1,3)} headers={post_headers} />
        <h2>{pokemon.name?.toUpperCase()}</h2>
        {pokemon.sprites?.front_default && (
          <img src={pokemon.sprites.front_default} alt={pokemon.name} />
        )}
      </div>
    </>
  );
}

export default App;
