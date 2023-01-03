import NewGameForm from "./NewGameForm";
import styles from "./new.module.scss";

export default function NewGame() {
  return (
    <div className={styles.container}>
      <NewGameForm />
    </div>
  );
}
