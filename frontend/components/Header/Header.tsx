import Link from "next/link";
import styles from "./Header.module.scss";
import cn from "../../utils/classnames";

const Header = () => (
  <div className={styles.container}>
    <Link href={"."} className={styles.link}>
      <h1 className={cn("title", styles.title)}>Minesweeper</h1>
    </Link>
  </div>
);

export default Header;
