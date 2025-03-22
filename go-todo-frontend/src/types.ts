export interface Todo {
  [key: string]: any; // Allow any key to be accessed with any value
  ID: number;
  Title: string;
  Completed: boolean;
}

export interface Post {
  [key: string]: any; // Allow any key to be accessed with any value
  id: number,
  title: string,
  body: string
}

interface Pokemon {
  id: number;
  name: string;
  base_experience: number;
  height: number;
  weight: number;
  abilities: Ability[];
  types: PokemonType[];
  sprites: Sprites;
  stats: Stat[];
  moves: Move[];
}

interface Ability {
  ability: {
    name: string;
    url: string;
  };
  is_hidden: boolean;
  slot: number;
}

interface PokemonType {
  slot: number;
  type: {
    name: string;
    url: string;
  };
}

interface Sprites {
  front_default: string;
  back_default: string;
  front_shiny: string;
  back_shiny: string;
  other: {
    official_artwork: {
      front_default: string;
    };
  };
}

interface Stat {
  base_stat: number;
  effort: number;
  stat: {
    name: string;
    url: string;
  };
}

interface Move {
  move: {
    name: string;
    url: string;
  };
}
