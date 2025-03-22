import React from "react"; // Import React
import { render, screen, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom";
import { vi } from "vitest";
import App from "../src/App";
import * as api from "../src/api";

vi.mock(import("../src/api"), async (importOriginal) => {
  const actual = await importOriginal();
  return {
    ...actual,
    // your mocked methods
    fetchTodos: vi.fn(() =>
      Promise.resolve([
        { ID: 1, Title: "Test Todo", Completed: false },
        { ID: 2, Title: "Another Todo", Completed: true },
      ])
    ),
  };
});

// vi.mock("../src/api", () => ({
//     fetchTodos: vi.fn(() =>
//         Promise.resolve([
//             { ID: 1, Title: "Test Todo", Completed: false },
//             { ID: 2, Title: "Another Todo", Completed: true },
//         ])
//     ),
// }));

describe("App Component", () => {
  test("renders the todo list correctly", async () => {
    render(<App />);
    expect(await screen.findByText("Test Todo")).toBeInTheDocument();
    expect(screen.getByText("Another Todo")).toBeInTheDocument();
    expect(screen.getByText("✅")).toBeInTheDocument();
    expect(screen.getByText("❌")).toBeInTheDocument();
  });
  test("increments count on button click", () => {
    render(<App />);
    const button = screen.getByText(/count is 0/i);
    fireEvent.click(button);
    expect(button).toHaveTextContent("count is 1");
  });
});
