import axios from "axios";
import { Todo, Post, Pokemon } from "./types";

const API_URL = "http://localhost:8080/todos";

export const fetchTodos = async (): Promise<Todo[]> => {
  const response = await axios.get(API_URL);
  console.log(response);

  return response.data;
};

export const fetchMe = async (): Promise<Post[]> => {
  const response = await axios.get(
    "https://jsonplaceholder.typicode.com/posts"
  );
  console.log(response);

  return response.data;
};

export const fetchPokemon = async (): Promise<Pokemon> => {
  const response = await fetch("https://pokeapi.co/api/v2/pokemon/charmander");
  console.log(response);
  return response.json();
};
