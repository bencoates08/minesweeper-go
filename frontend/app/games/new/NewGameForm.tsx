"use client";

import { useState } from "react";
import styles from "./NewGameForm.module.scss";
import cn from "../../../utils/classnames";
import newGame from "../../../apis/minesweeper-backend/newGame";
import { useRouter } from "next/navigation";

type NewGameForm = {
  name: string;
  height: string;
  width: string;
  bombs: string;
};

const defaultState: NewGameForm = {
  name: "",
  height: "",
  width: "",
  bombs: "",
};

export default function NewGameForm() {
  const router = useRouter();
  const [state, setState] = useState<NewGameForm>(defaultState);

  function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
    const value = event.target.value;
    setState({
      ...state,
      [event.target.name]: value,
    });
  }

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    // TODO; handle response
    const { name, height, width, bombs } = state;
    if (name && height && width && bombs) {
      newGame({
        name,
        height: parseInt(height),
        width: parseInt(width),
        bombs: parseInt(bombs),
      })
        .then((response) => {
          router.push(`/games/${response.id}`);
        })
        .catch((error) => {
          console.log(error);
          console.error(error);
        });
    }
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <input
        name="name"
        type="text"
        placeholder="Name"
        className={styles.input}
        required
        maxLength={30}
        value={state.name}
        onChange={handleChange}
      />
      <input
        name="height"
        type="number"
        placeholder="Height"
        className={styles.input}
        required
        max={99}
        value={state.height}
        onChange={handleChange}
      />
      <input
        name="width"
        type="number"
        placeholder="Width"
        className={styles.input}
        required
        max={99}
        value={state.width}
        onChange={handleChange}
      />
      <input
        name="bombs"
        type="number"
        placeholder="Bombs"
        className={styles.input}
        required
        max={
          state.height && state.width
            ? parseInt(state.height) * parseInt(state.width) - 1
            : 99
        }
        onChange={handleChange}
      />
      <input
        type="reset"
        value="Clear"
        className={cn(styles.buttonLeft, styles.transitionAll)}
        onClick={() => {
          setState(defaultState);
        }}
      />
      <input
        type="submit"
        value="Start Game"
        className={cn(styles.buttonRight, styles.transitionAll)}
      />
    </form>
  );
}
