import { useEffect, useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import { fetchTodos } from "./api";
import { Todo } from "./types";

function App() {
  const [count, setCount] = useState(0);
  const [todos, setTodos] = useState<Todo[]>([]);
  useEffect(() => {
    fetchTodos().then(setTodos);
  }, []);
  const headers = todos.length > 0 ? Object.keys(todos[0]) : [];
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
      </div>
      <div className="p-4">
        <h1 className="text-xl font-bold mb-4">Todo List</h1>
        {/* <ul>
          {todos.map((todo) => (
            <li key={todo.ID} className="p-2 border-b">
              {todo.Title} {todo.Completed ? "✅" : "❌"}
            </li>
          ))}
        </ul> */}
        <table>
          <thead>
            <tr>
              {headers.map((header, index) => (
                <th key={index}>{header}</th>
              ))}
            </tr>
          </thead>
          <tbody>
            {todos.map((todo, rowIndex) => (
              <tr key={rowIndex}>
                {headers.map((header, colIndex) => (
                  <td key={colIndex}>
                    {header === "Completed"
                      ? todo[header]
                        ? "✅"
                        : "❌"
                      : todo[header]}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
