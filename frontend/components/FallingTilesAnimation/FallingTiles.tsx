import styles from "./FallingTiles.module.scss";
import Image from "next/image";
import Tile1Img from "../../public/tile1.png";
import TileBombImg from "../../public/tilebomb.png";
import TileFlagImg from "../../public/tileflag.png";
import Tile3Img from "../../public/tile3.png";

const FallingTiles = () => (
  <>
    <Image
      className={styles.floatImg1}
      src={Tile1Img}
      alt="Minesweeper 1 tile"
      height={100}
      width={100}
    />
    <Image
      className={styles.floatImg2}
      src={TileBombImg}
      alt="Minesweeper bomb tile"
      height={100}
      width={100}
    />
    <Image
      className={styles.floatImg3}
      src={TileFlagImg}
      alt="Minesweeper flag tile"
      height={100}
      width={100}
    />
    <Image
      className={styles.floatImg4}
      src={Tile3Img}
      alt="Minesweeper 3 tile"
      height={100}
      width={100}
    />
    <Image
      className={styles.floatImg5}
      src={TileBombImg}
      alt="Minesweeper bomb tile"
      height={100}
      width={100}
    />
    <Image
      className={styles.floatImg6}
      src={TileFlagImg}
      alt="Minesweeper flag tile"
      height={100}
      width={100}
    />
  </>
);

export default FallingTiles;
