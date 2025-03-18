export interface Todo {
  [key: string]: any; // Allow any key to be accessed with any value
  ID: number;
  Title: string;
  Completed: boolean;
}
