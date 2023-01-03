import styles from "./home.module.scss";
import Link from "next/link";
import FallingTiles from "../components/FallingTilesAnimation/FallingTiles";
import cn from "../utils/classnames";

export default function HomePage() {
  return (
    <div className={cn(styles.homePage, styles.transitionAll)}>
      <FallingTiles />

      <h1 className={cn("title", styles.titleAnimation)}>Minesweeper</h1>
      <div className={cn(styles.controls, styles.inputAnimation)}>
        <Link
          key={Math.random()}
          href={"./games/new"}
          className={cn(styles.linkLeft, styles.transitionAll)}
        >
          New Game
        </Link>
        <Link
          key={Math.random()}
          href={"./games"}
          className={cn(styles.linkLeft, styles.transitionAll)}
        >
          My Games
        </Link>
      </div>
    </div>
  );
}
