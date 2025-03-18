import axios from "axios";
import { Todo } from "./types";

const API_URL = "http://localhost:8080/todos";

export const fetchTodos = async (): Promise<Todo[]> => {
  const response = await axios.get(API_URL);
  console.log(response);
  
  return response.data;
};
